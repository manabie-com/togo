package transport

import (
	"errors"
	"fmt"
	"net/http"
	"strings"

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
	logger *logs.Logger
	router *gin.Engine
	todo   *usecase.ToDoUsecase
}

func NewServer(postgres storages.Store) *Server {
	logger := logs.WithPrefix("Server")
	todo := usecase.NewToDoUsecase(postgres)

	server := &Server{
		logger: logger,
		todo:   todo,
	}

	server.setupRouter()
	return server
}

func (s *Server) setupRouter() {
	router := gin.Default()

	router.POST("/login", s.login)

	authGroups := router.Group("/").Use(s.authMiddleware())
	authGroups.POST("/tasks", s.addTask)
	authGroups.GET("/tasks/:created_date/:total/:page", s.listTasks)

	s.router = router
}

func (s *Server) Start(address string) error {
	return s.router.Run(address)
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

func successResponse(value interface{}) gin.H {
	return gin.H{"data": value}
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
