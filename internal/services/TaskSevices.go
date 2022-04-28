package services

import (
	"togo/internal/dao"
	"togo/internal/model"
	"togo/internal/utils"

	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
)

type TaskServices interface {
	CreateTask() error
}

func (c Con) CreateTask() error {

	payload := new(model.Task)
	utils.BodyParser(c.Ctx, payload)
	validate := validator.New()

	if err := validate.Struct(payload); err != nil {
		return c.Ctx.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	cdt := dao.CountDailyTask(payload.UserID, c.Db)
	mdt := dao.MaxDailyTodo(payload.UserID, c.Db)

	if cdt >= mdt {
		return c.Ctx.Status(fiber.StatusNotAcceptable).JSON(fiber.Map{"retCode": 406, "message": "Daily task limit exceeded"})
	}

	result := c.Db.Create(&payload)
	if result.Error != nil {
		return c.Ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"retCode": 500, "message": result.Error})

	}

	return c.Ctx.JSON(payload)
}
