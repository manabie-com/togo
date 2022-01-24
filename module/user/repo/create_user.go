package repo

import (
	"context"
	"gorm.io/gorm"
	"togo/module/user/model"
	model2 "togo/module/userconfig/model"
)

type CreateUserConfigStore interface {
	Create(ctx context.Context, data *model2.CreateUserConfig) error
}

type CreateUserStore interface {
	Create(ctx context.Context, data *model.CreateUser) error
	Get(ctx context.Context, cond map[string]interface{}) (*model.User, error)
}

type createUserRepo struct {
	store CreateUserStore
	userCfgStore CreateUserConfigStore
}

func NewCreateUserRepo(store CreateUserStore, userCfgStore CreateUserConfigStore) *createUserRepo {
	return &createUserRepo{store: store, userCfgStore: userCfgStore}
}

func (u *createUserRepo) CreateUser(ctx context.Context, data *model.CreateUser) error {
	cond := map[string]interface{}{
		"name": *data.Name,
	}

	if usr, err := u.store.Get(ctx, cond); usr != nil || err != gorm.ErrRecordNotFound {
		return err
	}

	if err := u.store.Create(ctx, data); err != nil {
		return err
	}

	userCgb := model2.CreateUserConfig{UserId: data.Id}
	if err := u.userCfgStore.Create(ctx, &userCgb); err != nil {
		return err
	}

	return nil
}