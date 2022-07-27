package handler

import (
	"togo/dto"
	"togo/pkg/logger"
	"togo/response"
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

func (h *UserHandler) Create(c echo.Context) error {
	createUserDto := new(dto.CreateUserDto)

	if errBindDto := c.Bind(createUserDto); errBindDto != nil {
		logger.L.Sugar().Errorf("[UserHandler] Create errBindDto: %s", errBindDto)
		return response.Error(c, errBindDto)
	}

	task, err := h.userService.Create(createUserDto)
	if err != nil {
		logger.L.Sugar().Errorf("[UserHandler] Create errCreateUser: %s", err)
		return response.Error(c, err)
	}
	return response.Success(c, task)
}
