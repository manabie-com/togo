package main

import (
	"togo-service/pkg/database"
	"togo-service/pkg/middleware"
	"togo-service/pkg/routes"
	"togo-service/pkg/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	app := fiber.New()

	middleware.FiberMiddleware(app)

	db := database.SetupDB()
	database.DoMigrate(db)
	routes.Setup(app, db)

	// Start server
	utils.StartServer(app)
}
