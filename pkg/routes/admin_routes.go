package routes

import (
	"togo-service/app/controllers"
	"togo-service/pkg/middleware"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

// PrivateRoutes func for describe group of private routes.
func AdminRoutes(a *fiber.App, db *gorm.DB) {
	handler := &controllers.Handler{
		DB: db,
	}

	// Create routes group.
	route := a.Group("/api/v1/admin")

	route.Post("/user-setting", middleware.JWTProtected(), handler.UpdateUser)
}
