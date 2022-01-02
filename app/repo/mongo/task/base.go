package task

import (
	"context"

	"github.com/manabie-com/togo/app/common/adapter"
	"github.com/manabie-com/togo/app/model"
)

type Repository interface {
	Create(ctx context.Context, in CreateReq) (model.Task, error)
	All(ctx context.Context, in AllReq) ([]model.Task, error)
}

type repoManager struct {
	db *adapter.CollectionV2
}

func NewRepoManager(db *adapter.CollectionV2) Repository {
	return &repoManager{db: db}
}
