package users

import (
	"fmt"
	"strconv"
	"todo/database"
	"todo/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

type UserController interface {
	Get(c *fiber.Ctx) error
	Create(c *fiber.Ctx) error
	Login(c *fiber.Ctx) error
}

type userController struct {
	responstory database.Responstory
}

type UserCreate struct {
	Name     string `json:"name"`
	Password string `json:"password"`
	Limit    uint   `json:"limit"`
}

type UserLogin struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

func InitUserController(responstory database.Responstory) UserController {
	controller := userController{responstory: responstory}
	return controller
}

func (control userController) Get(c *fiber.Ctx) error {
	id := c.Params("id")
	_, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Cannot find task",
		})
	}
	var task Users
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

func (control userController) Create(c *fiber.Ctx) error {

	var body UserCreate

	err := c.BodyParser(&body)

	if err != nil {
		fmt.Println(err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Cannot parse JSON",
		})
	}
	var exitsUser []Users
	if err := control.responstory.Find(&exitsUser, "name = ?", body.Name); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Server error",
		})
	}
	if len(exitsUser) > 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "User name exsits",
		})
	}

	hashedPassword, _ := utils.HashPassword(body.Password)
	user := Users{
		Name:     body.Name,
		Password: hashedPassword,
		Limit:    body.Limit,
	}

	if err := control.responstory.Insert(&user); err != nil {
		fmt.Println(err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Cannot insert data",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"success": true,
		"data": fiber.Map{
			"user": user,
		},
	})
}

func (control userController) Login(c *fiber.Ctx) error {
	var body UserLogin

	err := c.BodyParser(&body)

	var user Users

	if err := control.responstory.Find(&user, "name = ?", body.Name); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Cannot find user name",
		})
	}

	if match := utils.CheckPasswordHash(body.Password, user.Password); match != true {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Password not true",
		})
	}
	// Create the Claims
	claims := jwt.MapClaims{
		"name":  user.Name,
		"id":    user.Id,
		"limit": user.Limit,
		// "exp":   time.Now().Add(time.Hour * 72).Unix(),
	}

	// Create token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": "Server error",
		})
	}
	return c.JSON(fiber.Map{"token": t})
}
