package handler

import (
	"context"
	"errors"

	"github.com/google/uuid"

	"github.com/chi07/todo/internal/model"
)

var (
	ErrValidation = errors.New("validation error")
)

type CreateTaskService interface {
	CreateTask(ctx context.Context, t *model.Task) (uuid.UUID, error)
}
