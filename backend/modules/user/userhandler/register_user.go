package userhandler

import (
	"context"

	"github.com/phathdt/libs/go-sdk/sdkcm"
	"togo/common"
	"togo/modules/user/usermodel"
)

type RegisterUserRepo interface {
	FindUser(ctx context.Context, conditions map[string]interface{}) (*usermodel.User, error)
	CreateUser(ctx context.Context, data *usermodel.UserCreate) error
}

type Hasher interface {
	Hash(data string) string
}

type registerUserHdl struct {
	repo   RegisterUserRepo
	hasher Hasher
}

func NewRegisterUserHdl(repo RegisterUserRepo, hasher Hasher) *registerUserHdl {
	return &registerUserHdl{repo: repo, hasher: hasher}
}

func (h *registerUserHdl) Response(ctx context.Context, data *usermodel.UserCreate) error {
	user, _ := h.repo.FindUser(ctx, map[string]interface{}{"email": data.Email})

	if user != nil {
		return common.ErrEmailExisted
	}

	salt := common.GenSalt(50)

	data.Password = h.hasher.Hash(data.Password + salt)
	data.Salt = salt

	if err := h.repo.CreateUser(ctx, data); err != nil {
		return sdkcm.ErrCannotCreateEntity("user", err)
	}

	return nil
}
