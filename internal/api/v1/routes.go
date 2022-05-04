package v1

import (
	"github.com/gofiber/fiber/v2"
	"todo/configs"
	"todo/internal/api/middleware"
)

func MapRoutes(app *fiber.App, cfg *configs.Config, userHandler *UserHandler, taskHandler *TaskHandler) {
	v1 := app.Group("/api/v1")

	task := v1.Group("/tasks")
	task.Use(middleware.JWTMiddleware(cfg.JwtSecret))
	{
		task.Post("/", taskHandler.CreateTask)
		task.Get("/:id", taskHandler.GetTask)
		task.Get("/", taskHandler.GetTasks)
		task.Put("/:id", taskHandler.UpdateTask)
		task.Delete("/:id", taskHandler.DeleteTask)
	}

	v1.Post("/login", userHandler.Login)
	v1.Post("/signup", userHandler.SignUp)
}
