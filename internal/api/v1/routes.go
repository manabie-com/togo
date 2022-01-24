package v1

import "github.com/gofiber/fiber/v2"

func MapRoutes(app *fiber.App, userHandler *UserHandler, taskHandler *TaskHandler) {
	v1 := app.Group("/api/v1")
	{
		v1.Get("/tasks", taskHandler.CreateTask)
		v1.Post("/login", userHandler.Login)
	}
}
