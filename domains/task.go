package domains

import (
	"context"
	"time"
)

type (
	TaskRequest struct {
		CreatedDate time.Time
		UserId      int64
	}

	Task struct {
		Id          int64
		Content     string
		UserId      int64
		CreatedDate time.Time
	}

	TaskByIdRequest struct {
		Id     int64
		UserId int64
	}

	TaskRepository interface {
		GetCountCreatedTaskTodayByUser(ctx context.Context, userId int64) (int64, error)
		CreateTask(context.Context, *Task) (*Task, error)
		GetTasks(context.Context, *TaskRequest) ([]*Task, error)
		GetTaskById(context.Context, *TaskByIdRequest) (*Task, error)
	}
)
