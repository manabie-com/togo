package healthcheck

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func Init(group *echo.Group) {
	group.GET("", func(ctx echo.Context) error {
		return ctx.JSON(http.StatusOK, nil)
	})
}
