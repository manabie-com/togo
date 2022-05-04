package taskservice

import (
	"context"
	"github.com/sirupsen/logrus"
	"todo/internal/entities"
	"todo/pkg/helper"
)

func (t *taskService) CreateTask(ctx context.Context, task *entities.Task) (*entities.Task, error) {
	logrus.Info("CreateTask")
	user, err := helper.UserFromContext(ctx)
	if err != nil {
		logrus.Errorf("UserFromContext err %v", err)
		return nil, err
	}
	task.UserID = user.ID
	taskResp, err := t.taskRepo.Create(ctx, task)
	if err != nil {
		logrus.Errorf("Create task repo err %v", err)
		return nil, err
	}
	return taskResp, nil
}
