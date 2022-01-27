package taskservice

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	"github.com/trinhdaiphuc/togo/pkg/helper"
)

func (t *taskService) DeleteTask(ctx context.Context, id int) error {
	user, err := helper.UserFromContext(ctx)
	if err != nil {
		logrus.Errorf("UserFromContext err %v", err)
		return err
	}
	task, err := t.taskRepo.GetTask(ctx, id)
	if err != nil {
		return err
	}

	if task.UserID != user.ID {
		return fiber.ErrForbidden
	}
	return t.taskRepo.DeleteTask(ctx, id)
}
