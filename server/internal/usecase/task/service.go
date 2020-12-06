package task

import (
	"context"
	taskDB "github.com/HoangVyDuong/togo/internal/storages/task"
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

func (ts *taskService) GetTasks(ctx context.Context, userId string) ([]taskDB.Task, error) {
	return nil, nil
}

func (ts *taskService) CreateTask(ctx context.Context, task taskDB.Task) (string, error) {
	return "", nil
}