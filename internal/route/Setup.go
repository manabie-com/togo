package route

import "github.com/gofiber/fiber/v2"

func Setup(app *fiber.App) {
	v1 := app.Group("api/v1")
	TaskRoutes(v1)
	UserRoutes(v1)

}
