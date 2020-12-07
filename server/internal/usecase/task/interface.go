package task

import (
	"context"
	"github.com/HoangVyDuong/togo/internal/storages/task"
)

//Repository interface
type Repository interface {
	RetrieveTasks(ctx context.Context, userId int64) ([]task.Task, error)
	AddTask(ctx context.Context, taskEntity task.Task) (int64, error)
	SoftDeleteTask(ctx context.Context, taskId int64) error
}

//UseCase interface
type Service interface {
	GetTasks(ctx context.Context, userId int64) ([]task.Task, error)
	CreateTask(ctx context.Context, taskEntity task.Task) (int64, error)
	DeleteTask(ctx context.Context, taskId int64) error
}
