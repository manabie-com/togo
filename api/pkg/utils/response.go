package utils

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type APIError struct {
	Message string `json:"message"`
}

func ResponseSuccess(c echo.Context, i interface{}) error {
	return c.JSON(http.StatusOK, i)
}

func ResponseFailure(c echo.Context, code int, err error) error {
	return c.JSON(code, APIError{
		Message: err.Error(),
	})
}
