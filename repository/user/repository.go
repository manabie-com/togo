package user

import (
	"context"

	"github.com/khangjig/togo/model"
)

type Repository interface {
	GetByID(ctx context.Context, id int64) (*model.User, error)
	GetByEmail(ctx context.Context, email string) (*model.User, error)
}

type CacheRepository interface{}
