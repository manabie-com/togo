package repo

import (
	"context"
	"gorm.io/gorm"
	"togo/module/user/model"
)

type GetUserStore interface {
	Get(ctx context.Context, cond map[string]interface{}) (*model.User, error)
}

type getUserRepo struct {
	store GetUserStore
}

func NewGetUserRepo(store GetUserStore) *getUserRepo {
	return &getUserRepo{store: store}
}

func (u *getUserRepo) GetUser(ctx context.Context, cond map[string]interface{}) (*model.User, error) {
	usr, err := u.store.Get(ctx, cond)
	if err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return usr, nil
}
