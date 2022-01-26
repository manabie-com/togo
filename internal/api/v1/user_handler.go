package v1

import (
	"github.com/gofiber/fiber/v2"
	"github.com/trinhdaiphuc/togo/internal/entities"
	"github.com/trinhdaiphuc/togo/internal/service"
)

type UserHandler struct {
	service *service.Service
}

func NewUserHandler(service *service.Service) *UserHandler {
	return &UserHandler{service: service}
}

func (h *UserHandler) Login(ctx *fiber.Ctx) error {
	user := &entities.User{}
	err := ctx.BodyParser(user)
	if err != nil {
		return err
	}

	userResponse, err := h.service.UserService.Login(ctx.Context(), user)
	if err != nil {
		return err
	}
	return ctx.Status(fiber.StatusOK).JSON(userResponse)
}
