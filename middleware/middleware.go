package middleware

import (
	"errors"
	"strconv"
	"togo/constants"
	"togo/response"

	"github.com/labstack/echo/v4"
)

func Middleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			if err := next(c); err != nil {
				c.Error(err)
			}
			return nil
		}
	}
}

func TaskMiddlerware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			userID, errParse := strconv.Atoi(c.Param("id"))
			if errParse != nil {
				return response.Error(c, errors.New("missing user id"))
			}
			c.Set(string(constants.ContextUserID), userID)
			if err := next(c); err != nil {
				c.Error(err)
			}
			return nil
		}
	}
}
