package routes

import (
	"github.com/gin-gonic/gin"
	"togo/pkg/services"
)

func RegisterRoutes(r *gin.Engine, s *services.Server) {
	routes := r.Group("/auth")
	routes.POST("/register", s.RegisterAccount)
	routes.POST("/login", s.LoginAccount)

	routesTask := r.Group("/task")
	routesTask.Use(s.Validate)
	routesTask.GET("/list", s.GetTaskByDate)
	routesTask.POST("/create", s.CreateTask)
}
