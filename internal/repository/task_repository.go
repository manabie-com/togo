package repository

//go:generate mockgen -destination=./mock/mock_$GOFILE -source=$GOFILE -package=mock

import (
	"context"
	"github.com/trinhdaiphuc/togo/internal/entities"
	"github.com/trinhdaiphuc/togo/internal/infrastructure"
	"time"
)

type TaskRepository interface {
	Create(ctx context.Context, task *entities.Task) (*entities.Task, error)
}

type taskRepositoryImpl struct {
	db infrastructure.DB
}

func NewTaskRepository(db infrastructure.DB) TaskRepository {
	return &taskRepositoryImpl{
		db: db,
	}
}

func (t *taskRepositoryImpl) Create(ctx context.Context, task *entities.Task) (*entities.Task, error) {
	resp, err := t.db.Task.Create().
		SetName(task.Name).
		SetContent(task.Content).
		SetUserID(task.UserID).
		Save(ctx)
	if err != nil {
		return nil, err
	}
	return &entities.Task{
		ID:        resp.ID,
		Name:      resp.Name,
		Content:   resp.Content,
		UserID:    resp.UserID,
		CreatedAt: resp.CreatedAt.Format(time.RFC3339),
		UpdatedAt: resp.UpdatedAt.Format(time.RFC3339),
	}, nil
}
