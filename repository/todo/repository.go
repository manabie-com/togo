package todo

import (
	"context"

	"github.com/khangjig/togo/codetype"
	"github.com/khangjig/togo/model"
)

type Repository interface {
	Create(ctx context.Context, todo *model.Todo) error
	GetByID(ctx context.Context, id int64) (*model.Todo, error)
	Update(ctx context.Context, todo *model.Todo) error
	DeleteByID(ctx context.Context, id int64, unscoped bool) error

	GetList(
		ctx context.Context,
		userID int64,
		conditions interface{},
		search string,
		order string,
		paginator codetype.Paginator,
	) ([]model.Todo, int64, error)
}

type CacheRepository interface{}
