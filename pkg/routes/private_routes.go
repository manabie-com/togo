package routes

import (
	"togo-service/app/controllers"
	"togo-service/pkg/middleware"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func PrivateRoutes(a *fiber.App, db *gorm.DB) {
	handler := &controllers.Handler{
		DB: db,
	}

	// Create routes group.
	route := a.Group("/api/v1")

	route.Post("/task", middleware.JWTProtected(), handler.CreateTask)
	route.Put("/task/:task_id", middleware.JWTProtected(), handler.UpdateTask)
	route.Delete("/task/:task_id", middleware.JWTProtected(), handler.DeleteTask)
	route.Get("/tasks", middleware.JWTProtected(), handler.FetchTasks)
}
