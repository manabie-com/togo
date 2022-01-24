package handler

import (
	"context"
	"togo/module/userconfig/model"
)

type CreateUserConfigRepo interface {
	CreateUserConfig(ctx context.Context, data *model.CreateUserConfig) error
}

type createUserConfigHdl struct {
	userCfbCreateRepo CreateUserConfigRepo
}

func NewCreateUserConfigHdl(userCfbCreateRepo CreateUserConfigRepo) *createUserConfigHdl {
	return &createUserConfigHdl{userCfbCreateRepo: userCfbCreateRepo}
}

func (u *createUserConfigHdl) CreateUserConfig(ctx context.Context, data *model.CreateUserConfig) error {
	if err := u.userCfbCreateRepo.CreateUserConfig(ctx, data); err != nil {
		return err
	}

	return nil
}