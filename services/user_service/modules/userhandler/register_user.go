package userhandler

import (
	"context"

	"github.com/phathdt/libs/go-sdk/sdkcm"
	"user_service/common"
	"user_service/modules/usermodel"
)

type RegisterUserStorage interface {
	FindUser(ctx context.Context, conditions map[string]interface{}) (*usermodel.User, error)
	CreateUser(ctx context.Context, data *usermodel.UserCreate) error
}

type Hasher interface {
	Hash(data string) string
}

type registerUserHdl struct {
	storage RegisterUserStorage
	hasher  Hasher
}

func NewRegisterUserHdl(storage RegisterUserStorage, hasher Hasher) *registerUserHdl {
	return &registerUserHdl{storage: storage, hasher: hasher}
}

func (h *registerUserHdl) Response(ctx context.Context, data *usermodel.UserCreate) error {
	user, _ := h.storage.FindUser(ctx, map[string]interface{}{"email": data.Email})

	if user != nil {
		return common.ErrEmailExisted
	}

	salt := common.GenSalt(50)

	data.Password = h.hasher.Hash(data.Password + salt)
	data.Salt = salt

	if err := h.storage.CreateUser(ctx, data); err != nil {
		return sdkcm.ErrCannotCreateEntity("user", err)
	}

	return nil
}
