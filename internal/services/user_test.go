package services

import (
	"context"
	"github.com/manabie-com/togo/internal/utils"
	"github.com/stretchr/testify/assert"
	"testing"
)

type getStorageMock struct {}

var (
	isValidate func (ctx context.Context, id, password string) (bool, error)
	s *userService
)
func (s *getStorageMock ) ValidateUser(ctx context.Context, id, password string) (bool, error){
	return isValidate(ctx, id ,password)
}

func initUserService() {
	storage := &getStorageMock{}
	s = &userService{
		storage,
	}
}

func TestGetAuthToken_Valid(t *testing.T) {
	initUserService()
	isValidate = func (ctx context.Context, id, password string) (bool, error) {
		return true, nil
	}
	res, err := s.GetAuthToken(context.TODO(),"1", "123456")
	assert.NotNil(t, res)
	assert.Nil(t, err)
	assert.EqualValues(t, true, res)
}

func TestGetAuthToken_NotValidID(t *testing.T) {
	initUserService()
	isValidate = func (ctx context.Context, id, password string) (bool, error) {
		return false, &utils.NotFoundError{}
	}
	res, err := s.GetAuthToken(context.TODO(),"1", "123456")
	assert.NotNil(t, err)
	assert.EqualValues(t, false, res)
}


