package main

import (
	"fmt"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/joho/godotenv"
	"github.com/manabie-com/togo/common/response"
	"github.com/manabie-com/togo/db"
	"github.com/manabie-com/togo/middlewares"
	"github.com/manabie-com/togo/modules/auth"
	"github.com/manabie-com/togo/modules/tasks"
	"github.com/manabie-com/togo/modules/users"
	postgre "github.com/manabie-com/togo/repo/postgres"
)

var services map[string]interface{}

func init() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Cannot load env file")
	}
	db.ConnectPostgre()
	services = initServices()
}

func main() {
	port := os.Getenv("PORT")
	app := fiber.New(fiber.Config{
		ReadBufferSize: 8192, // Prevent header too large error of long access token
	})

	app.Use(cors.New())
	app.Use(recover.New())

	// Health check
	app.Get("/health-check", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(response.SuccessResponse{
			Status:  fiber.StatusOK,
			Data:    "I'm fine!",
			Message: "OK",
		})
	})

	app.Use(middlewares.NewAuthenticator(services["authService"].(auth.Service)))

	// Add start route below
	tasks.StartTasksRoute(app, services)
	users.StartUsersRoute(app, services)

	// Not found handle
	app.Use(func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusNotFound).JSON(response.ErrorResponse{
			Status:  fiber.StatusNotFound,
			Error:   "not_found",
			Message: "URL not found, please check URL again.",
		})
	})

	fmt.Println("App is running and listening on port " + port)
	app.Listen(":" + port)
}

func initServices() map[string]interface{} {
	fmt.Println("Initialing app services")

	// Init Repositories
	taskRepo := postgre.TasksRepo{}
	userRepo := postgre.UsersRepo{}

	// Init services
	authService := auth.Service{
		JWTKey: "wqGyEBBfPK9w3Lxw",
	}
	tasksService := tasks.Service{
		Repo: taskRepo,
	}
	usersService := users.Service{
		Repo: userRepo,
	}

	fmt.Println("App services has been initialized")

	return map[string]interface{}{
		"authService":  authService,
		"tasksService": tasksService,
		"usersService": usersService,
	}
}
