package userbiz

import (
	"context"
	"github.com/japananh/togo/common"
	"github.com/japananh/togo/modules/user/usermodel"
)

type RegisterStore interface {
	FindUser(ctx context.Context, conditions map[string]interface{}, moreInfo ...string) (*usermodel.User, error)
	CreateUser(ctx context.Context, data *usermodel.UserCreate) error
}

type Hasher interface {
	Hash(data string) string
}

type registerBiz struct {
	store  RegisterStore
	hasher Hasher
}

func NewRegisterBiz(store RegisterStore, hasher Hasher) *registerBiz {
	return &registerBiz{store: store, hasher: hasher}
}

func (biz *registerBiz) Register(ctx context.Context, data *usermodel.UserCreate) error {
	if err := data.Validate(); err != nil {
		return err
	}

	user, err := biz.store.FindUser(ctx, map[string]interface{}{"email": data.Email})

	if user != nil {
		return common.ErrEntityExisted(usermodel.EntityName, err)
	}

	if err == common.ErrRecordNotFound {
		salt := common.GenSalt(50)

		data.Password = biz.hasher.Hash(data.Password + salt)
		data.Salt = salt

		if err := biz.store.CreateUser(ctx, data); err != nil {
			return common.ErrCannotCreateEntity(usermodel.EntityName, err)
		}
	} else {
		return common.ErrDB(err)
	}

	return nil
}
