package util

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/khangjig/togo/util/myerror"
)

type response struct{}

var Response response

func (response) Success(c echo.Context, data interface{}) error {
	return c.JSON(http.StatusOK, data)
}

func (response) Error(c echo.Context, err myerror.MyError) error {
	var errMessage string

	if err.Raw != nil {
		errMessage = err.Raw.Error()
	}

	return c.JSON(err.HTTPCode, map[string]interface{}{
		"code":    err.ErrorCode,
		"message": err.Message,
		"error":   errMessage,
	})
}
