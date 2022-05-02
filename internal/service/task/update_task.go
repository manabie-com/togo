package taskservice

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	"todo/internal/entities"
	"todo/pkg/helper"
)

func (t *taskService) UpdateTask(ctx context.Context, task *entities.Task) (*entities.Task, error) {
	user, err := helper.UserFromContext(ctx)
	if err != nil {
		logrus.Errorf("UserFromContext err %v", err)
		return nil, err
	}
	resp, err := t.taskRepo.GetTask(ctx, task.ID)
	if err != nil {
		return nil, err
	}

	if resp.UserID != user.ID {
		return nil, fiber.ErrForbidden
	}
	return t.taskRepo.UpdateTask(ctx, task)
}
