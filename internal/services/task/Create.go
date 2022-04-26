package task

import (
	"togo/internal/connect"
	"togo/internal/model"
	"togo/internal/services/user"
	"togo/internal/utils"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func Create(c *fiber.Ctx) error {

	payload := new(model.Task)
	utils.BodyParser(c, payload)
	validate := validator.New()

	if err := validate.Struct(payload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	count := CountDaily(payload.UserID)
	max := user.GetMaxDailyTodo(payload.UserID)

	if count > max {
		return c.Status(fiber.StatusNotAcceptable).JSON(fiber.Map{"retCode": 400, "message": "Already reached the maximum daily task."})
	}

	result := connect.DB.Create(&payload)
	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"retCode": 500, "message": result.Error})

	}

	return c.JSON(payload)
}
