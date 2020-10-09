package user

import (
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/manabie-com/togo/internal/middleware"
)

// LoadRoute func
func LoadRoute(e *echo.Echo, controller Controller) {
	fmt.Println("Load route user")
	g := e.Group("/user")
	g.POST("/register", controller.Register)

	g.GET("/me", controller.GetUser, middleware.IsAuthenticate)
}
