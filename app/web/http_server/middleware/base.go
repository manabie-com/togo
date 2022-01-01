package middleware

import (
	"github.com/labstack/echo/v4"
	"github.com/manabie-com/togo/app/common/config"

	gHandler "github.com/manabie-com/togo/app/common/gstuff/handler"
)

var cfg = config.GetConfig()

func getRequestID(c echo.Context) (requestID string) {
	return gHandler.GetRequestID(c)
}

// CallFromPublicRoute set "from_public_route":true, detect user call api via public route
func CallFromPublicRoute() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			c.Set("from_public_route", true) // bool
			return next(c)
		}
	}
}
