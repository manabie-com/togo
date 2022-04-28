package route

import (
	"togo/internal/controller"
	"togo/internal/middleware"

	"github.com/gofiber/fiber/v2"
)

func TaskRoutes(rg fiber.Router) {

	tg := rg.Group("/task")
	tg.Post("/create", middleware.Protected(), controller.CreateTask)

}
