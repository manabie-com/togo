package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/luongdn/togo/config"
	"github.com/luongdn/togo/database"
	"github.com/luongdn/togo/routes"
)

func main() {
	config.Load()
	database.ConnectMySQL()
	database.ConnectRedis()

	// Create some test data
	database.Seed()

	app := fiber.New()
	routes.SetUpRoutes(app)

	log.Fatal(app.Listen(":8080"))
}
