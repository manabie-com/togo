package service

import (
	"context"
	"log"
	"togo/domain/errdef"
	"togo/domain/model"
	"togo/domain/repository"
)

type TaskService interface {
	CreateTask(ctx context.Context, user *model.User, task *model.Task) (*model.Task, error)
}

type taskServiceImpl struct {
	repo repository.TaskRepository
}

func (this *taskServiceImpl) CreateTask(ctx context.Context, user *model.User, task *model.Task) (*model.Task, error) {
	countTaskCreatedInDay, err := this.repo.CountTaskCreatedInDayByUser(ctx, *user)
	log.Printf("taskInDay: %d", countTaskCreatedInDay)
	if countTaskCreatedInDay >= user.Limit {
		return nil, errdef.LimitTaskCreated
	}
	task.CreatedBy = user.Id
	err = this.repo.Create(ctx, *task)
	return task, err
}

func NewTaskService(repo repository.TaskRepository) TaskService {
	return &taskServiceImpl{
		repo: repo,
	}
}
