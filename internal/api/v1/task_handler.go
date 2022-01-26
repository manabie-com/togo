package v1

import (
	"github.com/gofiber/fiber/v2"
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
	task := &entities.Task{}
	err := ctx.BodyParser(task)
	if err != nil {
		return err
	}

	taskResponse, err := t.service.TaskService.CreateTask(ctx.Context(), task)
	if err != nil {
		return err
	}
	return ctx.Status(fiber.StatusCreated).JSON(taskResponse)
}
