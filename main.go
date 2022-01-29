package main

import (
	"fmt"
	"todo/database"
	"todo/modules/tasks"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	mysqlClient, _ := database.Initmysql(&tasks.Tasks{})
	tasksresponstory := database.InitResponstory(mysqlClient)
	tasksController := tasks.InitTaskController(tasksresponstory)
	fmt.Println(mysqlClient)
	app.Get("/", tasksController.Get)
	app.Post("/", tasksController.Create)

	app.Listen(":5000")
}
