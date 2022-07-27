package controllers

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
)

func (ctrl *Controller) loadMux() {
	e := echo.New()

	e.Use(middleware.RateLimiter(middleware.NewRateLimiterMemoryStore(1000)))
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		Skipper:          middleware.DefaultSkipper,
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{http.MethodGet, http.MethodHead, http.MethodPut, http.MethodPatch, http.MethodPost, http.MethodDelete},
		AllowCredentials: true,
	}))

	apiv1 := e.Group("/api")
	togov1 := apiv1.Group("/togo/v1")
	{
		task := togov1.Group("/task")
		{
			task.GET("", apiGetTask(ctrl))
			task.POST("", apiPostTask(ctrl))
		}
	}
	{
		user := togov1.Group("/user")
		{
			user.POST("/register", apiPostUser(ctrl))
		}
	}
	ctrl.e = e

}
