package response

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func Error(c echo.Context, err error) error {
	return c.JSON(http.StatusOK, echo.Map{"status": 400, "data": err.Error()})
}

func Success(c echo.Context, msg interface{}) error {
	return c.JSON(http.StatusOK, echo.Map{"status": http.StatusOK, "data": msg})
}
