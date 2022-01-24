package handler

import (
	"context"
	"togo/module/user/model"
)

type CreateUserRepo interface {
	CreateUser(ctx context.Context, data *model.CreateUser) error
}

type createUserHdl struct {
	taskCreateRepo CreateUserRepo
}

func NewCreateUserHdl(taskCreateRepo CreateUserRepo) *createUserHdl {
	return &createUserHdl{taskCreateRepo: taskCreateRepo}
}

func (u *createUserHdl) CreateUser(ctx context.Context, data *model.CreateUser) error {
	if err := u.taskCreateRepo.CreateUser(ctx, data); err != nil {
		return err
	}

	return nil
}