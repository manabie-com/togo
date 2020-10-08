package task

import (
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/manabie-com/togo/internal/middleware"
)

// LoadRoute func
func LoadRoute(e *echo.Echo, controller Controller) {
	fmt.Println("Load route task")
	g := e.Group("/tasks")

	g.Use(middleware.IsAuthenticate)

	g.POST("", controller.AddTask)
	g.GET("", controller.RetrieveTasks)

	g.POST("/many", controller.AddManyTasks)
}
