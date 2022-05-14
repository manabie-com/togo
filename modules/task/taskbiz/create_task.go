package taskbiz

import (
	"context"
	"github.com/japananh/togo/modules/task/taskmodel"
)

type createTaskRepo interface {
	CreateTask(ctx context.Context, data *taskmodel.TaskCreate) error
}

type createTaskBiz struct {
	repo createTaskRepo
}

func NewCreateTaskBiz(repo createTaskRepo) *createTaskBiz {
	return &createTaskBiz{repo: repo}
}

func (biz *createTaskBiz) CreateTask(ctx context.Context, data *taskmodel.TaskCreate) error {
	if err := biz.repo.CreateTask(ctx, data); err != nil {
		return err
	}

	return nil
}
