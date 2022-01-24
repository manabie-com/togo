package repo

import (
	"context"
	"togo/module/userconfig/model"
)

type UpdateUserConfigStore interface {
	Update(ctx context.Context, cond map[string]interface{}, data *model.UpdateUserConfig) error
	Get(ctx context.Context, cond map[string]interface{}) (*model.UserConfig, error)
}

type updateUserConfigRepo struct {
	store UpdateUserConfigStore
}

func NewUpdateUserConfigRepo(store UpdateUserConfigStore) *updateUserConfigRepo {
	return &updateUserConfigRepo{store: store}
}

func (u *updateUserConfigRepo) UpdateUserConfig(ctx context.Context, userId uint, data *model.UpdateUserConfig) error {
	cond := map[string]interface{}{
		"user_id": userId,
	}

	_, err := u.store.Get(ctx, cond)
	if err != nil {
		return err
	}

	if err := u.store.Update(ctx, cond, data); err != nil {
		return err
	}

	return nil
}