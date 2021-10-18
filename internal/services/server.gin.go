package services

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/jericogantuangco/togo/internal/storages/postgres"
	"github.com/jericogantuangco/togo/internal/token"
	"github.com/jericogantuangco/togo/internal/util"
)



type Server struct {
	TokenMaker token.Maker
	Store  postgres.Store
	Router *gin.Engine
}

type ResponseMessage struct {
	Message string `json:"message"`
}

func NewServer(store postgres.Store) (*Server, error) {
	tokenMaker, err := token.NewJWTMaker(util.RandomString(32))
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %w", err)
	}
	server := &Server{
		Store: store,
		TokenMaker: tokenMaker,
	}
	router := gin.Default()

	router.GET("/login", server.login)

	authRoutes := router.Group("/").Use(authMiddleware(server.TokenMaker))
	authRoutes.GET("/tasks", server.listTasks)
	authRoutes.POST("/tasks", server.addTask)

	server.Router = router
	return server, nil
}

func (server *Server) Start(address string) error {
	return server.Router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{
		"error": err.Error(),
	}
}
