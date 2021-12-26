package storage

import (
	"context"
	"errors"
	"time"
	"togo-internal-service/internal/model"
)

var (
	ErrTaskNotFound          = errors.New("task not found")
	ErrExceedLimitTaskPerDay = errors.New("exceed limit task created per day")
	ErrInternal              = errors.New("internal server error")
)

type StorageConfig struct {
	MaxTaskCreatedPerDay int
	SubstrContentLength  int
	SonyflakeStartTime   time.Time
}

type Storage interface {
	ListTask(ctx context.Context, userID string, date time.Time, limit, offset int) ([]*model.Task, error)
	GetTask(ctx context.Context, ID string) (*model.Task, error)
	CreateTask(ctx context.Context, task *model.Task) (*model.Task, error)
	Close() error
}
