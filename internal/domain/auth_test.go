package domain

import (
	"context"
	"fmt"
	"testing"

	"manabie/togo/common/errors"
	"manabie/togo/internal/model"
	"manabie/togo/utils"

	"github.com/bxcodec/faker/v3"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func Test_Auth_Login_UserIDNotExist(t *testing.T) {
	userStore := &mockUserStore{}
	jwtKey := "someKey"
	authDomain := NewAuthDomain(userStore, jwtKey)

	userStore.On("FindUser", mock.Anything).Return(nil, fmt.Errorf("something error"))
	ctx := context.Background()
	token, err := authDomain.Login(ctx, &model.User{
		ID: faker.Username(),
	})

	assert.Equal(t, errors.ErrUserDoesNotExist, err)
	assert.Empty(t, token)

}

func Test_Auth_Login_WrongPassword(t *testing.T) {
	userStore := &mockUserStore{}
	jwtKey := "someKey"
	userID := faker.Username()
	password := faker.Password()
	authDomain := NewAuthDomain(userStore, jwtKey)

	userStore.On("FindUser", mock.Anything).Return(&model.User{
		Password: "wrong password",
	}, nil)
	ctx := context.Background()
	token, err := authDomain.Login(ctx, &model.User{
		ID:       userID,
		Password: password,
	})

	assert.Error(t, err)
	assert.Empty(t, token)

}

func Test_Auth_Login_Success(t *testing.T) {
	userStore := &mockUserStore{}
	jwtKey := "someKey"
	userID := faker.Username()
	password := faker.Password()
	authDomain := NewAuthDomain(userStore, jwtKey)

	hashedPassword, err := utils.HashPassword(password)
	assert.NoError(t, err)
	userStore.On("FindUser", mock.Anything).Return(&model.User{
		Password: hashedPassword,
	}, nil)
	ctx := context.Background()
	token, err := authDomain.Login(ctx, &model.User{
		ID:       userID,
		Password: password,
	})

	assert.NoError(t, err)
	assert.NotEmpty(t, token)

}

func Test_Auth_Register_UserIDExist(t *testing.T) {
	userStore := &mockUserStore{}
	jwtKey := "someKey"
	authDomain := NewAuthDomain(userStore, jwtKey)

	userStore.On("FindUser", mock.Anything).Return(&model.User{}, nil)
	ctx := context.Background()
	err := authDomain.Register(ctx, &model.User{
		ID: faker.Username(),
	})

	assert.Equal(t, errors.ErruserAlreadyExist, err)

}

func Test_Auth_Register_CreateUserFail(t *testing.T) {
	userStore := &mockUserStore{}
	jwtKey := "someKey"
	authDomain := NewAuthDomain(userStore, jwtKey)

	rErr := fmt.Errorf("create failed")
	userStore.On("FindUser", mock.Anything).Return(nil, fmt.Errorf("something error"))
	userStore.On("Create", mock.Anything).Return(rErr)

	ctx := context.Background()
	err := authDomain.Register(ctx, &model.User{
		ID:       faker.Username(),
		Password: faker.Password(),
	})

	assert.Equal(t, rErr, err)

}

func Test_Auth_Register_Success(t *testing.T) {
	userStore := &mockUserStore{}
	jwtKey := "someKey"
	authDomain := NewAuthDomain(userStore, jwtKey)

	userStore.On("FindUser", mock.Anything).Return(nil, fmt.Errorf("something error"))
	userStore.On("Create", mock.Anything).Return(nil)

	ctx := context.Background()
	err := authDomain.Register(ctx, &model.User{
		ID:       "userID",
		Password: "password",
	})

	assert.NoError(t, err)

}
