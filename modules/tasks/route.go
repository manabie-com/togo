package tasks

import (
	"github.com/gofiber/fiber/v2"
)

func StartTasksRoute(app *fiber.App, services map[string]interface{}) {
	controllers := TasksController{
		Srv: services["tasksService"].(Service),
	}
	group := app.Group("/tasks")

	group.Post("/", controllers.Create)
	group.Get("/", controllers.GetList)
}
