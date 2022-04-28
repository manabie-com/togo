package route

import (
	"togo/internal/controller"

	"github.com/gofiber/fiber/v2"
)

func UserRoutes(rg fiber.Router) {

	tg := rg.Group("/user")
	tg.Post("/create", controller.CreateUser)
	tg.Get("/login", controller.LoginUser)

}
