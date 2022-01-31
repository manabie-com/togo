package tasks

import (
	"fmt"
	"strconv"
	"time"
	"todo/database"
	"todo/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
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
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userid := uint(claims["id"].(float64))
	userlimit := uint(claims["limit"].(float64))

	fmt.Println(userid, userlimit)

	id := c.Params("id")
	_, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Cannot find task",
		})
	}
	var task Tasks
	if err := control.responstory.Get(&task, id); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Cannot find task",
		})
	}
	fmt.Println(task)
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"data": fiber.Map{
			"task": task,
		},
	})
}

func (control taskController) Create(c *fiber.Ctx) error {
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userid := strconv.Itoa(int(claims["id"].(float64)))
	userlimit := int(claims["limit"].(float64))

	fmt.Println(userid, userlimit)
	now := time.Now()
	start := strconv.Itoa(int(utils.StartOfDay(now).Unix()))
	end := strconv.Itoa(int(utils.EndOfDay(now).Unix()))
	var exitsTasks []Tasks
	if err := control.responstory.Find(&exitsTasks, "created_by = ? AND created_at BETWEEN ? AND ?", userid, start, end); err != nil {
		fmt.Println(err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Cannot insert data",
		})
	}
	if len(exitsTasks) >= int(userlimit) {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "User had task more then limmit perday",
		})
	}
	var body TasksCreate

	err := c.BodyParser(&body)

	if err != nil {
		fmt.Println(err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Cannot parse JSON",
		})
	}

	task := Tasks{
		Title:      body.Title,
		Desciption: body.Discription,
		CreatedBy:  uint(claims["id"].(float64)),
	}

	if err := control.responstory.Insert(&task); err != nil {
		fmt.Println(err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Cannot insert data",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"success": true,
		"data": fiber.Map{
			"task": task,
		},
	})
}
