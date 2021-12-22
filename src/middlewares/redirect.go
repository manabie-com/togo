package middlewares

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func HTTPSRedirect() echo.MiddlewareFunc {
	return middleware.HTTPSRedirectWithConfig(middleware.RedirectConfig{
		Code: http.StatusMovedPermanently,
	})
}

func NonWWWRedirect() echo.MiddlewareFunc {
	return middleware.NonWWWRedirectWithConfig(middleware.RedirectConfig{
		Code: http.StatusMovedPermanently,
	})
}
