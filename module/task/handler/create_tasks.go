package handler

import (
	"context"
	model2 "togo/module/task/model"
)

type CreateTaskRepo interface {
	CreateTasks(ctx context.Context, userId uint, data []model2.CreateTask) error
}

type createTaskHdl struct {
	taskCreateRepo CreateTaskRepo
}

func NewCreateTaskHdl(taskCreateRepo CreateTaskRepo) *createTaskHdl {
	return &createTaskHdl{taskCreateRepo: taskCreateRepo}
}

func (u *createTaskHdl) CreateTasks(ctx context.Context, userId uint, data []model2.CreateTask) error {
	if err := u.taskCreateRepo.CreateTasks(ctx, userId, data); err != nil {
		return err
	}

	return nil
}