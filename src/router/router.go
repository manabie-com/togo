package router

import (
	"net/http"
	"todo-api/service"
	"todo-api/src/controllers"
	"todo-api/src/errors"
	"todo-api/src/middlewares"
	"todo-api/src/models"

	"github.com/labstack/echo/v4"
)

func Init(app *echo.Echo, s *service.Service) {
	app.Use(middlewares.Cors())
	app.Use(middlewares.Gzip())
	app.Use(middlewares.Logger())
	app.Use(middlewares.Secure())
	app.Use(middlewares.Recover())
	app.HTTPErrorHandler = errors.HttpErrorHandler

	app.GET("/", controllers.Index())

	api := app.Group("/api")
	{
		task := api.Group("/task", MockAuthMiddleware())
		{
			task.GET("", controllers.GetTasks(s))
			task.POST("", controllers.CreateTask(s))
			task.DELETE("/:id", controllers.DeleteTask(s))
		}
	}
}

func MockAuthMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) (err error) {
			userID := c.Request().Header.Get("UserID")
			if userID == "" {
				return c.JSON(http.StatusForbidden, models.BaseResponse{
					Message: "Unauthorized",
				})
			}
			c.Set("UserID", userID)
			return next(c)
		}
	}
}
