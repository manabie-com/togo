package main

import (
	"todo/database"
	"todo/modules/tasks"
	"todo/modules/users"

	"github.com/gofiber/fiber/v2"

	jwtware "github.com/gofiber/jwt/v3"
)

func main() {
	app := fiber.New()

	mysqlClient, _ := database.Initmysql(&tasks.Tasks{}, &users.Users{})

	tasksresponstory := database.InitResponstory(mysqlClient, "tasks")
	tasksController := tasks.InitTaskController(tasksresponstory)

	usersresponstory := database.InitResponstory(mysqlClient, "users")
	usersController := users.InitUserController(usersresponstory)

	app.Get("/user/:id", usersController.Get)
	app.Post("/user/", usersController.Create)
	app.Post("/user/login", usersController.Login)

	app.Use(jwtware.New(jwtware.Config{
		SigningKey: []byte("secret"),
	}))

	app.Get("/:id", tasksController.Get)
	app.Post("/", tasksController.Create)

	app.Listen(":5000")
}
