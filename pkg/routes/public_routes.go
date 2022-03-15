package routes

import (
	"togo-service/app/controllers"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

// PublicRoutes func for describe group of public routes.
func PublicRoutes(a *fiber.App, db *gorm.DB) {
	handler := &controllers.Handler{
		DB: db,
	}

	a.Get("/", func(app *fiber.Ctx) error {
		return app.JSON(map[string]string{
			"message": "Togo API",
			"version": "v1",
		})
	})
	// Create routes group.
	route := a.Group("/api/v1")

	route.Post("/register", handler.Register)
	route.Post("/login", handler.Login)
}
