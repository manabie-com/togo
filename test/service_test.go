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
	s          usecase.UserService
)

func (s *getRespositoryMock) ValidateUser(ctx context.Context, id, password string) (bool, error) {
	return isValidate(ctx, id, password)
}

func initUserService() {
	respository := &getRespositoryMock{}
	s = usecase.NewUserService(respository)
}

func TestGetAuthToken_Valid(t *testing.T) {
	initUserService()
	isValidate = func(ctx context.Context, id, password string) (bool, error) {
		return true, nil
	}
	res, err := s.GetAuthToken(context.TODO(), "test", "123456")
	assert.NotNil(t, res)
	assert.Nil(t, err)
	assert.EqualValues(t, true, res)
}

func TestGetAuthToken_NotValidID(t *testing.T) {
	initUserService()
	isValidate = func(ctx context.Context, id, password string) (bool, error) {
		return false, &utils.NotFoundError{}
	}
	res, err := s.GetAuthToken(context.TODO(), "test", "123456")
	assert.NotNil(t, err)
	assert.EqualValues(t, false, res)
}
