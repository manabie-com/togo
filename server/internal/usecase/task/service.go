package task

import (
	"context"
	taskEntity "github.com/HoangVyDuong/togo/internal/storages/task"
)

type taskService struct {
	repo Repository
}

//NewService create new service
func NewService(r Repository) Service {
	return &taskService{
		repo: r,
	}
}


func (ts *taskService) GetTasks(ctx context.Context, userId int64) ([]taskEntity.Task, error) {
	return nil, nil
}

func (ts *taskService) CreateTask(ctx context.Context, taskEntity taskEntity.Task) (int64, error) {
	return 0, nil
}

func (ts *taskService) DeleteTask(ctx context.Context, taskId int64) error {
	return nil
}