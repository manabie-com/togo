package v1

import (
	"strconv"

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

func (t *TaskHandler) GetTask(ctx *fiber.Ctx) error {
	id, err := strconv.ParseInt(ctx.Params("id"), 10, 64)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid id")
	}
	taskResponse, err := t.service.TaskService.GetTask(ctx.Context(), int(id))
	if err != nil {
		return err
	}
	return ctx.JSON(taskResponse)
}

func (t *TaskHandler) GetTasks(ctx *fiber.Ctx) error {
	page, err := strconv.ParseInt(ctx.Query("page"), 10, 64)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid page param")
	}
	limit, err := strconv.ParseInt(ctx.Query("limit"), 10, 64)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid limit param")
	}
	filter := &entities.TaskFilter{
		Page:  int(page),
		Limit: int(limit),
	}
	tasksResponse, err := t.service.TaskService.GetTasks(ctx.Context(), filter)
	if err != nil {
		return err
	}
	return ctx.JSON(tasksResponse)
}

func (t *TaskHandler) UpdateTask(ctx *fiber.Ctx) error {
	id, err := strconv.ParseInt(ctx.Params("id"), 10, 64)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid id")
	}
	task := &entities.Task{
		ID: int(id),
	}

	if err := ctx.BodyParser(task); err != nil {
		return err
	}

	taskResponse, err := t.service.TaskService.UpdateTask(ctx.Context(), task)
	if err != nil {
		return err
	}
	return ctx.JSON(taskResponse)
}

func (t *TaskHandler) DeleteTask(ctx *fiber.Ctx) error {
	id, err := strconv.ParseInt(ctx.Params("id"), 10, 64)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid id")
	}
	err = t.service.TaskService.DeleteTask(ctx.Context(), int(id))
	if err != nil {
		return err
	}
	return ctx.SendStatus(fiber.StatusNoContent)
}
