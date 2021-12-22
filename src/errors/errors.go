package errors

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type httpError struct {
	code    int
	Key     string `json:"error"`
	Message string `json:"message"`
}

func NewHTTPError(code int, key string, msg string) *httpError {
	return &httpError{
		code:    code,
		Key:     key,
		Message: msg,
	}
}

func (e *httpError) Error() string {
	return e.Key + ": " + e.Message
}

func HttpErrorHandler(err error, c echo.Context) {
	var (
		code = http.StatusInternalServerError
		key  = "ServerError"
		msg  string
	)

	if he, ok := err.(*httpError); ok {
		code = he.code
		key = he.Key
		msg = he.Message
	} else if true {
		msg = err.Error()
	} else {
		msg = http.StatusText(code)
	}

	if !c.Response().Committed {
		if c.Request().Method == echo.HEAD {
			err := c.NoContent(code)
			if err != nil {
				c.Logger().Error(err)
			}
		} else {
			err := c.JSON(code, NewHTTPError(code, key, msg))
			if err != nil {
				c.Logger().Error(err)
			}
		}
	}
}
