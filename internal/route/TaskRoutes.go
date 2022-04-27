package route

import (
	"togo/internal/middleware"
	"togo/internal/services/task"

	"github.com/gofiber/fiber/v2"
)

func TaskRoutes(rg fiber.Router) {

	tg := rg.Group("/task")
	tg.Post("/create", middleware.Protected(), task.Create)

}
