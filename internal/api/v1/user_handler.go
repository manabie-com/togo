package v1

import (
	"github.com/gofiber/fiber/v2"
	"github.com/trinhdaiphuc/logger"
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
	var (
		user = &entities.User{}
		log  = logger.GetLogger(ctx.Context())
	)
	err := ctx.BodyParser(user)
	if err != nil {
		log.Error(err)
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request body",
			"error":   err,
		})
	}

	userResponse, err := h.service.UserService.Login(ctx, user)
	if err != nil {
		log.Error(err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Internal server error",
			"error":   err,
		})
	}
	return ctx.Status(fiber.StatusOK).JSON(userResponse)
}
