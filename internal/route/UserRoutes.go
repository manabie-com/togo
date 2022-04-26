package route

import (
	"togo/internal/services/user"

	"github.com/gofiber/fiber/v2"
)

func UserRoutes(rg fiber.Router) {

	tg := rg.Group("/user")
	tg.Post("/create", user.Create)

	// rg.GET("/user/:id", user.Get)
	// rg.GET("/user", user.GetAll)
	// rg.PUT("/user/:id", user.Update)
	// rg.DELETE("/user/:id", user.Delete)

}
