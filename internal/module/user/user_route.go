package user

import (
	"fmt"

	"github.com/labstack/echo/v4"
)

// LoadRoute func
func LoadRoute(e *echo.Echo, controller Controller) {
	fmt.Println("Load route user")
	g := e.Group("/user")
	g.POST("/register", controller.Register)
}
