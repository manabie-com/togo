package task

import (
	"context"
	"github.com/HoangVyDuong/togo/internal/storages/task"
)

//Repository interface
type Repository interface {
	RetrieveTasks(ctx context.Context, userId int64) ([]task.Task, error)
	AddTask(ctx context.Context, task task.Task) error
}

//UseCase interface
type Service interface {
	GetTasks(ctx context.Context, userId string) ([]task.Task, error)
	CreateTask(ctx context.Context, task task.Task) (string, error)
}
