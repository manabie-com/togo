package service

import (
	"testing"
	"togo/internal/models"
	"togo/internal/user/dto"
	"togo/internal/user/mocks"

	"github.com/test-go/testify/assert"
)

func mockCreateUserDtoAndUser() (*dto.CreateUserDto, *models.User) {
	createUserDto := &dto.CreateUserDto{
		Name:       "name",
		LimitCount: 1,
	}
	user := &models.User{
		Name:       createUserDto.Name,
		LimitCount: createUserDto.LimitCount,
	}
	return createUserDto, user
}

func TestUserService_CreateUserSuccess(t *testing.T) {
	createUserDto, user := mockCreateUserDtoAndUser()

	repo := mocks.NewUserRepository(t)
	repo.On("GetByName", createUserDto.Name).Return(nil, nil)
	repo.On("Create", user).Return(user, nil)

	service := NewUserService(repo)

	userResponse, err := service.Create(createUserDto)
	assert.Nil(t, err)
	assert.Equal(t, user.Name, userResponse.Name)
	assert.Equal(t, user.LimitCount, userResponse.LimitCount)
}

func TestUserService_CannotCreateUserWithNameExist(t *testing.T) {
	createUserDto, user := mockCreateUserDtoAndUser()

	repo := mocks.NewUserRepository(t)
	repo.On("GetByName", createUserDto.Name).Return(user, nil)

	service := NewUserService(repo)
	service.Create(createUserDto)

	u, err := service.Create(createUserDto)
	assert.Nil(t, u)
	assert.Equal(t, err.Error(), "user_with_name_exist")
}
