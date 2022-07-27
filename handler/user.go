package handler

import (
	"togo/service"

	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	userService *service.UserService
}

func NewUserHandler(
	router *echo.Group,
	userService *service.UserService,
) {
	handler := &UserHandler{userService}
	router.POST("", handler.Create)
}

func (h *UserHandler) Create(e echo.Context) error {
	return nil
}
