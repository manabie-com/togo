package users

import (
	"github.com/gofiber/fiber/v2"
)

func StartUsersRoute(app *fiber.App, services map[string]interface{}) {
	controllers := UsersController{
		Srv: services["usersService"].(Service),
	}
	group := app.Group("/users")

	group.Get("/login", controllers.Login)
	app.Get("/login", controllers.Login)
}
