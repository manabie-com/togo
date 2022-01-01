package route

import (
	"github.com/labstack/echo/v4"

	"github.com/manabie-com/togo/app/common/gstuff/handler"
	gmiddleware "github.com/manabie-com/togo/app/common/gstuff/middleware"
)

func baseRoute(e *echo.Echo) *echo.Group {
	return e.Group("", gmiddleware.LogBody)
}

// APIRoute ..
func APIRoute(e *echo.Echo) *echo.Group {
	base := baseRoute(e)
	base.Any("", handler.Health)
	apiRoute := base.Group("/api")
	apiRoute.Any("/health", handler.Health)
	return apiRoute
}

// PublicAPIRoute ..
func PublicAPIRoute(e *echo.Echo) *echo.Group {
	base := baseRoute(e)
	base.Any("", handler.Health)
	apiRoute := base.Group("/public-api")
	apiRoute.Any("/health", handler.Health)
	return apiRoute
}
