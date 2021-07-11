package task

import (
	"context"
	"time"

	"github.com/google/uuid"

	"github.com/manabie-com/togo/internal/entity"
	"github.com/manabie-com/togo/internal/storages"
	"github.com/manabie-com/togo/pkg/server/middleware"
)

type (
	service struct {
		taskStorage storages.TaskStorage
	}
	Service interface {
		List(ctx context.Context, createdDate string) ([]*entity.Task, error)
		Create(ctx context.Context, content string) (*entity.Task, error)
	}
)

func NewTaskService(taskStorage storages.TaskStorage) Service {
	return &service{
		taskStorage: taskStorage,
	}
}

func (s *service) List(ctx context.Context, createdDate string) ([]*entity.Task, error) {
	userID, _ := middleware.UserIDFromCtx(ctx)
	return s.taskStorage.RetrieveTasks(ctx, userID, createdDate)
}

func (s *service) Create(ctx context.Context, content string) (*entity.Task, error) {
	t := &entity.Task{}
	now := time.Now()
	userID, _ := middleware.UserIDFromCtx(ctx)
	t.ID = uuid.New().String()
	t.UserID = userID
	t.Content = content
	t.CreatedDate = now.Format("2006-01-02")
	return t, s.taskStorage.AddTask(ctx, t)
}
