package handler

import (
	"context"
	"togo/module/userconfig/model"
)

type UpdateUserConfigRepo interface {
	UpdateUserConfig(ctx context.Context, userId uint, data *model.UpdateUserConfig) error
}

type updateUserConfigHdl struct {
	userCfgUpdateRepo UpdateUserConfigRepo
}

func NewUpdateUserConfigHdl(userCfgUpdateRepo UpdateUserConfigRepo) *updateUserConfigHdl {
	return &updateUserConfigHdl{userCfgUpdateRepo: userCfgUpdateRepo}
}

func (u *updateUserConfigHdl) UpdateUserConfig(ctx context.Context, userId uint, data *model.UpdateUserConfig) error {
	if err := u.userCfgUpdateRepo.UpdateUserConfig(ctx, userId, data); err != nil {
		return err
	}

	return nil
}