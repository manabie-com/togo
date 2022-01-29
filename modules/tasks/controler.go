package tasks

import (
	"fmt"
	"todo/database"

	"github.com/gofiber/fiber/v2"
)

type TaskController interface {
	Get(c *fiber.Ctx) error
	Create(c *fiber.Ctx) error
}

type taskController struct {
	responstory database.Responstory
}

type TasksCreate struct {
	Title       string `json:"title"`
	Discription string `json:"description"`
}

func InitTaskController(responstory database.Responstory) TaskController {
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

func (control taskController) Create(c *fiber.Ctx) error {

	var body TasksCreate

	err := c.BodyParser(&body)

	if err != nil {
		fmt.Println(err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Cannot parse JSON",
		})
	}

	task := &Tasks{
		Title:      body.Title,
		Desciption: body.Discription,
	}

	if err := control.responstory.Insert(task); err != nil {
		fmt.Println(err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Cannot insert data",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"success": true,
		"data": fiber.Map{
			"todo": task,
		},
	})
}
