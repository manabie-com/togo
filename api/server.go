package api

import (
	"fmt"
	"net/http"
	db "togo/db/sqlc"
	"togo/token"
	"togo/util"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

type Server struct {
	config     util.Config
	store      db.Store
	tokenMaker token.Maker
	router     *gin.Engine
}

func NewServer(config util.Config, store db.Store) (*Server, error) {
	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %w", err)
	}
	server := &Server{
		config:     config,
		store:      store,
		tokenMaker: tokenMaker,
	}
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("fullname", validFullname)
	}
	server.setupRouter()
	return server, nil
}

func (server *Server) setupRouter() {
	router := gin.Default()
	router.POST("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"text": "togo"})
	})
	router.POST("/users", server.createUser)
	router.POST("/users/login", server.loginUser)
	// Add authentication middleware, using Paseto Token
	authRoutes := router.Group("/").Use(authMiddleware(server.tokenMaker))
	// These endpoints need authorization, implements in their handlers repsectively
	authRoutes.GET("/users", server.listUsers)
	authRoutes.GET("/tasks", server.listTasks)
	authRoutes.POST("/tasks", server.createTask)
	server.router = router
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
