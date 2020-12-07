package task

import (
	"context"
	"github.com/HoangVyDuong/togo/internal/storages/task"
	"time"
)

//Repository interface
type Repository interface {
	RetrieveTasks(ctx context.Context, userId uint64) ([]task.Task, error)
	AddTask(ctx context.Context, taskEntity task.Task, createdAt time.Time) error
	SoftDeleteTask(ctx context.Context, taskId uint64, deletedAt time.Time) error
}

//UseCase interface
type Service interface {
	GetTasks(ctx context.Context, userId uint64) ([]task.Task, error)
	CreateTask(ctx context.Context, taskEntity task.Task) error
	DeleteTask(ctx context.Context, taskId uint64) error
}
