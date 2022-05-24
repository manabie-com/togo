package userbiz_test

import (
	"context"
	"github.com/japananh/togo/common"
	"github.com/japananh/togo/modules/user/userbiz"
	"github.com/japananh/togo/modules/user/usermodel"
	"github.com/stretchr/testify/assert"
	"testing"
)

type mockRegisterStore struct{}

func (mockRegisterStore) FindUser(_ context.Context, conditions map[string]interface{}, _ ...string) (*usermodel.User, error) {
	if val, ok := conditions["email"]; ok && val == "user@gmail.com" {
		return &usermodel.User{Email: conditions["email"].(string)}, nil
	}
	return nil, common.ErrRecordNotFound
}

func (mockRegisterStore) CreateUser(_ context.Context, data *usermodel.UserCreate) error {
	data.Id = 2
	return nil
}

type mockHash struct{}

func (mockHash) Hash(data string) string {
	return data
}

func TestUserBiz_RegisterSuccess(t *testing.T) {
	biz := userbiz.NewRegisterBiz(mockRegisterStore{}, mockHash{})
	err := biz.Register(nil, &usermodel.UserCreate{Email: "user1@gmail.com", Password: "user@123"})
	assert.Nil(t, err)
}

func TestUserBiz_RegisterErrEmailExisted(t *testing.T) {
	biz := userbiz.NewRegisterBiz(mockRegisterStore{}, mockHash{})
	err := biz.Register(nil, &usermodel.UserCreate{Email: "user@gmail.com", Password: "user@123"})
	assert.NotNil(t, err)
}
