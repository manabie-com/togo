package transport

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/manabie-com/togo/internal/logs"
	"github.com/manabie-com/togo/internal/storages/postgres"
	"github.com/manabie-com/togo/internal/usecase"
	"github.com/manabie-com/togo/internal/util"
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

func NewServer() *Server {
	logger := logs.WithPrefix("Server")
	postgres := postgres.NewPostgres()
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

	authGroups := router.Group("/").Use(authMiddleware())
	authGroups.POST("/tasks", s.addTask)
	authGroups.GET("/tasks/:created_date/:total/:page", s.listTasks)

	s.router = router
}

func (s *Server) Start(address string) error {
	return s.router.Run(address)
}

func authMiddleware() gin.HandlerFunc {
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

		id, status := validToken(fields[1])
		if !status {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, errorResponse(errors.New("token is invalid")))
			return
		}

		ctx.Set(authorizationPayloadKey, id)
		ctx.Next()
	}
}

func validToken(token string) (string, bool) {
	claims := make(jwt.MapClaims)
	t, err := jwt.ParseWithClaims(token, claims, func(*jwt.Token) (interface{}, error) {
		return []byte(util.Conf.SecretKey), nil
	})
	if err != nil {
		log.Println(err)
		return "", false
	}

	if !t.Valid {
		return "", false
	}

	id, ok := claims["user_id"].(string)
	if !ok {
		return "", false
	}

	return id, true
}

func successResponse(value interface{}) gin.H {
	return gin.H{"data": value}
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
