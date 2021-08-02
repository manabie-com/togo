package rest

import (
	"github.com/gofiber/fiber/v2"
	"togo/internal/entity"
	"togo/utils/validator"
)

func SimpleError(c *fiber.Ctx, err error) error {
	resp := validator.ToErrResponse(err)

	if resp == nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"errors": err.Error(),
		})
	}

	return c.Status(fiber.StatusInternalServerError).JSON(resp)
}

func GetCurrentUser(c *fiber.Ctx) *entity.User {
	currentUser := c.Locals("currentUser").(*entity.User)

	return currentUser
}
