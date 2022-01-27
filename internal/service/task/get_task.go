package taskservice

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	"github.com/trinhdaiphuc/togo/internal/entities"
	"github.com/trinhdaiphuc/togo/pkg/helper"
)

func (t *taskService) GetTask(ctx context.Context, id int) (*entities.Task, error) {
	user, err := helper.UserFromContext(ctx)
	if err != nil {
		logrus.Errorf("UserFromContext err %v", err)
		return nil, err
	}
	task, err := t.taskRepo.GetTask(ctx, id)
	if err != nil {
		return nil, err
	}

	if task.UserID != user.ID {
		return nil, fiber.ErrForbidden
	}
	return task, nil
}
