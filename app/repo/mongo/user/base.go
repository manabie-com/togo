package user

import (
	"context"

	"github.com/manabie-com/togo/app/common/adapter"
	"github.com/manabie-com/togo/app/model"
)

type Repository interface {
	Create(ctx context.Context, in CreateReq) (model.User, error)
	GetOneByID(ctx context.Context, id int) (model.User, error)
	GetOneByUsername(ctx context.Context, username string) (model.User, error)
	IncNumTask(ctx context.Context, in IncNumTaskReq) (model.User, error)
}

type repoManager struct {
	db *adapter.CollectionV2
}

func NewRepoManager(db *adapter.CollectionV2) Repository {
	return &repoManager{db: db}
}
