package controllers

import (
	"github.com/labstack/echo/v4"
)

func Index() echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.JSON(200, "Welcome to Echo")
	}
}
