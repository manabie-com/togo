package taskservice

import (
	"context"
	"github.com/sirupsen/logrus"
	"github.com/trinhdaiphuc/togo/internal/entities"
	"github.com/trinhdaiphuc/togo/pkg/helper"
)

func (t *taskService) GetTasks(ctx context.Context, filter *entities.TaskFilter) (*entities.Tasks, error) {
	logrus.Info("GetTasks")
	user, err := helper.UserFromContext(ctx)
	if err != nil {
		logrus.Errorf("UserFromContext err %v", err)
		return nil, err
	}
	filter.UserID = user.ID
	return t.taskRepo.GetTasks(ctx, filter)
}
