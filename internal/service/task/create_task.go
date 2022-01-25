package taskservice

import (
	"github.com/gofiber/fiber/v2"
	"github.com/trinhdaiphuc/togo/internal/entities"
	"github.com/trinhdaiphuc/togo/pkg/helper"
)

func (t *taskService) CreateTask(ctx *fiber.Ctx, task *entities.Task) (*entities.Task, error) {
	user, err := helper.UserFromContext(ctx)
	if err != nil {
		return nil, err
	}
	task.UserID = user.ID
	taskResp, err := t.taskRepo.Create(ctx.Context(), task)
	if err != nil {
		return nil, err
	}
	return taskResp, nil
}
