package handler

import (
	"togo/internal/common"
	"togo/internal/response"
	"togo/internal/user/dto"
	"togo/internal/user/service"

	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	userService service.UserService
}

func NewUserHandler(
	router *echo.Group,
	userService service.UserService,
) *UserHandler {
	handler := &UserHandler{userService}
	router.POST("", handler.Create)

	return handler
}

func (h *UserHandler) Create(c echo.Context) error {
	createUserDto := new(dto.CreateUserDto)

	if errBindValidate := common.BindValidate(c, createUserDto); errBindValidate != nil {
		//logger.L.Sugar().Errorf("[UserHandler] Create errBindValidate: %s", errBindValidate)
		return response.Error(c, errBindValidate)
	}

	task, err := h.userService.Create(createUserDto)
	if err != nil {
		//logger.L.Sugar().Errorf("[UserHandler] Create errCreateUser: %s", err)
		return response.Error(c, err)
	}
	return response.Success(c, task)
}
