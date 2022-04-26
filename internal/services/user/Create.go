package user

import (
	"togo/internal/connect"
	"togo/internal/model"
	"togo/internal/utils"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func Create(c *fiber.Ctx) error {

	payload := new(model.User)
	utils.BodyParser(c, payload)
	validate := validator.New()

	if err := validate.Struct(payload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	hash, err := utils.HashPassword(payload.Password)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	payload.Password = hash

	result := connect.DB.Create(&payload)
	if result.Error != nil {
		return c.Status(500).JSON(fiber.Map{"retCode": 500, "message": result.Error})

	}

	return c.JSON(payload)
}
