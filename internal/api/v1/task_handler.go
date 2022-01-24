package v1

import (
	"github.com/gofiber/fiber/v2"
	"github.com/trinhdaiphuc/logger"
	"github.com/trinhdaiphuc/togo/internal/entities"
	"github.com/trinhdaiphuc/togo/internal/service"
)

type TaskHandler struct {
	service *service.Service
}

func NewTaskHandler(service *service.Service) *TaskHandler {
	return &TaskHandler{service: service}
}

func (t *TaskHandler) CreateTask(ctx *fiber.Ctx) error {
	var (
		task = &entities.Task{}
		log  = logger.GetLogger(ctx.Context())
	)
	err := ctx.BodyParser(task)
	if err != nil {
		log.Error(err)
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request body",
			"err":     err,
		})
	}

	taskResponse, err := t.service.TaskService.CreateTask(ctx, task)
	if err != nil {
		log.Error(err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Internal server error",
			"err":     err,
		})
	}
	return ctx.Status(fiber.StatusOK).JSON(taskResponse)
}
