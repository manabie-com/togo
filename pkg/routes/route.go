package routes

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func Setup(app *fiber.App, db *gorm.DB) {
	PublicRoutes(app, db)
	PrivateRoutes(app, db)
	AdminRoutes(app, db)
	NotFoundRoute(app)
}
