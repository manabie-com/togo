package user

import (
	"context"

	"github.com/manabie-com/togo/app/common/adapter"
	"github.com/manabie-com/togo/app/model"
)

type Repository interface {
	Create(ctx context.Context, in CreateReq) (model.User, error)
	// Update(ctx context.Context, req UpdateReq) (out model.Property, err error)
	// GetOneByID(ctx context.Context, id primitive.ObjectID) (out model.Property, err error)
	// All(ctx context.Context, in AllReq) (out []model.Property, err error)
}

type repoManager struct {
	db *adapter.CollectionV2
}

func NewRepoManager(db *adapter.CollectionV2) Repository {
	return &repoManager{db: db}
}
