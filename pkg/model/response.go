package model

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

type Response struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
	Error   interface{} `json:"error"`
}

func ResponseSuccess(c echo.Context, data interface{}) error {
	return c.JSON(http.StatusOK, Response{
		Status:  http.StatusOK,
		Message: "Success",
		Data:    data,
	})
}

func ResponseWithError(c echo.Context, err interface{}) error {
	return c.JSON(http.StatusOK, Response{
		Status:  http.StatusBadRequest,
		Message: "Error",
		Error:   err,
	})
}
