package grpc

import (
	"context"

	"github.com/khangjig/togo/model"
)

type IUseCase interface {
	GetTodoByID(ctx context.Context, todoID int64) (*model.Todo, error)
}
