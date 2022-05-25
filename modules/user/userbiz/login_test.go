package userbiz_test

import (
	"context"
	"github.com/japananh/togo/common"
	"github.com/japananh/togo/component/tokenprovider"
	"github.com/japananh/togo/modules/user/userbiz"
	"github.com/japananh/togo/modules/user/usermodel"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

type mockLoginStore struct{}

func (mockLoginStore) FindUser(_ context.Context, conditions map[string]interface{}, _ ...string) (*usermodel.User, error) {
	if val, ok := conditions["email"]; ok && val.(string) == "user@gmail.com" {
		return &usermodel.User{Email: val.(string), Password: "user@123", Salt: ""}, nil
	}
	return nil, common.ErrRecordNotFound
}

type mockProvider struct{}

func (mockProvider) Generate(_ tokenprovider.TokenPayload, expiry int) (*tokenprovider.Token, error) {
	return &tokenprovider.Token{
		Token:   "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJwYXlsb2FkIjp7InVzZXJfaWQiOjF9LCJleHAiOjE2NTM0MDk1MDksImlhdCI6MTY1MzMyMzEwOX0.SYzR9JXyIc_VeuLXLAnxFWTM3nO6LQfWbyO-vTK3fMo",
		Expiry:  expiry,
		Created: time.Now().UTC(),
	}, nil
}

func (mockProvider) Validate(_ string) (*tokenprovider.TokenPayload, error) {
	return &tokenprovider.TokenPayload{}, nil
}

func TestLoginBiz_LoginSucceed(t *testing.T) {
	biz := userbiz.NewLoginBiz(
		mockLoginStore{},
		mockProvider{},
		mockHash{},
		&tokenprovider.TokenConfig{AccessTokenExpiry: 86400, RefreshTokenExpiry: 604800},
	)
	user, err := biz.Login(nil, &usermodel.UserLogin{Email: "user@gmail.com", Password: "user@123"})
	assert.Nil(t, err)
	assert.NotNil(t, user)
}

func TestLoginBiz_LoginErr(t *testing.T) {
	biz := userbiz.NewLoginBiz(
		mockLoginStore{},
		mockProvider{},
		mockHash{},
		&tokenprovider.TokenConfig{AccessTokenExpiry: 86400, RefreshTokenExpiry: 604800},
	)
	user, err := biz.Login(nil, &usermodel.UserLogin{Email: "user1@gmail.com"})
	assert.NotNil(t, err)
	assert.Nil(t, user)
}
