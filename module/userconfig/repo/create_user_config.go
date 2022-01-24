package repo

import (
	"context"
	"errors"
	"togo/module/userconfig/model"
)

type CreateUserConfigStore interface {
	Create(ctx context.Context, data *model.CreateUserConfig) error
	Get(ctx context.Context, cond map[string]interface{}) (*model.UserConfig, error)
}

type createUserConfigRepo struct {
	store CreateUserConfigStore
}

func NewCreateUserConfigRepo(store CreateUserConfigStore) *createUserConfigRepo {
	return &createUserConfigRepo{store: store}
}

func (u *createUserConfigRepo) CreateUserConfig(ctx context.Context, data *model.CreateUserConfig) error {
	cond := map[string]interface{}{
		"user_id": *data.UserId,
	}

	usr, err := u.store.Get(ctx, cond)
	if err != nil {
		return err
	}

	if usr != nil {
		return errors.New("UserIsExisted")
	}

	if err := u.store.Create(ctx, data); err != nil {
		return err
	}

	return nil
}