package api

import (
	"github.com/gin-gonic/gin"
)

func setupRouter(server *Server) {
	router := gin.Default()
	server.router = router
	router.POST("/users", server.createUser)
	router.POST("/users/login", server.loginUser)
	authRoutes := router.Group("/").Use(authMiddleware(server.tokenMaker))
	authRoutes.GET("/users", server.getListUsers)
	authRoutes.POST("/tasks", server.createTask)
}
