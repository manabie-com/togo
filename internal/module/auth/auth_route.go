package auth

import (
	"fmt"

	"github.com/labstack/echo/v4"
)

// LoadRoute func
func LoadRoute(e *echo.Echo, controller Controller) {
	fmt.Println("Load route auth")

	e.POST("/login", controller.Login)
}
