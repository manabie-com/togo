package services

import (
	"context"
	"github.com/manabie-com/togo/internal/model"
)


type TaskService interface {
	ListTasks(ctx context.Context, userId string, createdDate string) (res []*model.Task, err error)
	AddTask(ctx context.Context, t *model.Task) error
	IsAllowedToAddTask(ctx context.Context, userId string) bool
}

type taskService struct {
	taskStorage model.TaskStorage
}

func NewTaskService(ts model.TaskStorage) TaskService{
	return &taskService{
		taskStorage: ts,
	}
}


func (s *taskService) ListTasks(ctx context.Context, userId string, createdDate string) (res []*model.Task, err error) {
	tasks, err := s.taskStorage.RetrieveTasks(
		ctx, userId, createdDate,
	)
	if err != nil {
		return nil, err
	}
	return tasks, err
}

func (s *taskService) AddTask(ctx context.Context, t *model.Task) error{

	err := s.taskStorage.AddTask(ctx, t)
	if err != nil {
		return err
	}
	return nil
}

func (s *taskService) IsAllowedToAddTask(ctx context.Context, userId string) bool {
	return s.taskStorage.IsAllowedToAddTask(ctx, userId)
}







