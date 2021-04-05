package transport

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/manabie-com/togo/internal/logs"
	"github.com/manabie-com/togo/internal/storages"

	"github.com/manabie-com/togo/internal/usecase"
)

const (
	authorizationHeaderKey  = "authorization"
	authorizationTypeBearer = "bearer"
	authorizationPayloadKey = "authorization_payload"
)

type Server struct {
	logger     *logs.Logger
	httpServer *http.Server
	todo       *usecase.ToDoUsecase
}

func NewServer(postgres storages.Store) *Server {
	logger := logs.WithPrefix("Server")
	todo := usecase.NewToDoUsecase(postgres)
	httpServer := &http.Server{}

	server := &Server{
		logger:     logger,
		todo:       todo,
		httpServer: httpServer,
	}

	server.setupRouter()

	return server
}

func (s *Server) setupRouter() {
	router := gin.Default()
	rateLimitGroups := router.Group("/").Use(s.rateLimitMiddleware())
	rateLimitGroups.POST("/login", s.login)

	authGroups := router.Group("/").Use(s.rateLimitMiddleware()).Use(s.authMiddleware())
	authGroups.POST("/tasks", s.addTask)
	authGroups.GET("/tasks/:created_date/:total/:page", s.listTasks)

	s.httpServer.Handler = router
}

func (s *Server) Start(address string) error {
	s.httpServer.Addr = address
	return s.httpServer.ListenAndServe()
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}

func (s *Server) authMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authorizationHeader := ctx.GetHeader(authorizationHeaderKey)

		if len(authorizationHeader) == 0 {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, errorResponse(errors.New("authorization is not provided")))
			return
		}

		fields := strings.Fields(authorizationHeader)
		if len(fields) < 2 {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, errorResponse(errors.New("invalid authorization header format")))
			return
		}

		authorizationType := strings.ToLower(fields[0])

		if authorizationType != authorizationTypeBearer {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, errorResponse(fmt.Errorf("unsupported authorization type %s", authorizationType)))
			return
		}

		id, status := s.todo.ValidToken(fields[1])
		if !status {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, errorResponse(errors.New("token is invalid")))
			return
		}

		ctx.Set(authorizationPayloadKey, id)
		ctx.Next()
	}
}

func (s *Server) rateLimitMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		if cip := c.ClientIP(); cip != "" {
			now := time.Now()
			limit := limiter.GetIP(cip)
			if limit.Count == 0 {
				limit := Limit{
					Count:    1,
					Interval: now,
				}
				limiter.AddIP(cip, limit)
				return
			} else {
				interval := now.Sub(limit.Interval).Seconds()
				if interval >= limiter.intervalSecond {
					limit.Count = 1
					limit.Interval = now
					limiter.AddIP(cip, limit)
					return
				}

				limit.Count = limit.Count + 1
				limit.Interval = now
				limiter.AddIP(cip, limit)
				if limit.Count >= limiter.maxRequest {
					c.AbortWithStatusJSON(http.StatusTooManyRequests, errorResponse(errors.New("too many request")))
					return
				}
			}
		}

		c.Next()
	}
}

func successResponse(value interface{}) gin.H {
	return gin.H{"data": value}
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
