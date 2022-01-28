package tasks

import (
	"github.com/gofiber/fiber/v2",
	"togo/models",
)

type TaskController interface {
	Get(c *fiber.Ctx) error
}

type taskController struct {
	responstory models.Responstory
}

func InitTaskController(responstory Responstory) TaskController {
	controller := taskController{responstory: responstory}
	return controller
}

func (control taskController) Get(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"data": fiber.Map{
			"todos": "",
		},
	})
}
