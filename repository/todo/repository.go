package todo

import (
	"context"

	"github.com/khangjig/togo/model"
)

type Repository interface {
	Create(ctx context.Context, channel *model.Todo) error
	GetByID(ctx context.Context, id int64) (*model.Todo, error)
	Update(ctx context.Context, channel *model.Todo) error
	Delete(ctx context.Context, channel *model.Todo) error
}

type CacheRepository interface{}
