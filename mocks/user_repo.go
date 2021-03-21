package mocks

import (
	"github.com/manabie-com/togo/models"
	"github.com/stretchr/testify/mock"
)

type IUserRepository struct {
	mock.Mock
}

func (mock *IUserRepository) GetUserByUserName(username string) (*models.User, error) {
	args := mock.Called(username)
	return args.Get(0).(*models.User), args.Error(1)
}

func (mock *IUserRepository) AddUser(user *models.User) (*models.User, error) {
	args := mock.Called(user)
	return args.Get(0).(*models.User), args.Error(1)
}
