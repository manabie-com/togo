package test

import (
	"context"
	"testing"

	"github.com/manabie-com/togo/internal/usecase"
	"github.com/manabie-com/togo/internal/utils"
	"github.com/stretchr/testify/assert"
)

type getRespositoryMock struct{}

var (
	isValidate func(ctx context.Context, id, password string) (bool, error)
	us         usecase.UserService
)

func (s *getRespositoryMock) ValidateUser(ctx context.Context, id, password string) (bool, error) {
	return isValidate(ctx, id, password)
}

func initUserService() {
	respository := &getRespositoryMock{}
	us = usecase.NewUserService(respository)
}

func TestGetAuthTokenWithValidID(t *testing.T) {
	initUserService()
	isValidate = func(ctx context.Context, id, password string) (bool, error) {
		return true, nil
	}
	res, err := us.GetAuthToken(context.TODO(), "test", "123456")
	assert.NotNil(t, res)
	assert.Nil(t, err)
	assert.EqualValues(t, true, res)
}

func TestGetAuthTokenWithInvalidID(t *testing.T) {
	initUserService()
	isValidate = func(ctx context.Context, id, password string) (bool, error) {
		return false, &utils.NotFoundError{}
	}
	res, err := us.GetAuthToken(context.TODO(), "test", "123456")
	assert.NotNil(t, err)
	assert.EqualValues(t, false, res)
}
