package apiutils

import (
	"net/http"

	"manabie/todo/models"

	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
)

var (
	ErrForbidden     = errors.New("access denied")
	ErrInvalidValue  = errors.New("invalid value")
	ErrIncorrectData = errors.New("incorrect data")
	ErrNotFound      = errors.New("not found")
)

func ResponseSuccess(c echo.Context, i interface{}) error {
	return c.JSON(http.StatusOK, i)
}

func ResponseFailure(c echo.Context, err error) error {
	code := http.StatusInternalServerError
	cause := errors.Cause(err)

	switch cause {
	case ErrForbidden:
		code = http.StatusForbidden
	case ErrInvalidValue:
		code = http.StatusBadRequest
	case ErrIncorrectData:
		code = http.StatusServiceUnavailable
	case ErrNotFound:
		code = http.StatusNotFound
	}

	return c.JSON(code, models.APIError{
		Message: err.Error(),
	})
}
