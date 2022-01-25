package v1

import (
	"github.com/gofiber/fiber/v2"
	"github.com/trinhdaiphuc/togo/configs"
	"github.com/trinhdaiphuc/togo/internal/api/middleware"
)

func MapRoutes(app *fiber.App, cfg *configs.Config, userHandler *UserHandler, taskHandler *TaskHandler) {
	v1 := app.Group("/api/v1")

	task := v1.Group("/tasks")
	task.Use(middleware.JWTMiddleware(cfg.JwtSecret))
	{
		task.Post("/", taskHandler.CreateTask)
	}

	v1.Post("/login", userHandler.Login)
}
