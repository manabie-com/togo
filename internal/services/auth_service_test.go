package services

import (
	"context"
	"github.com/manabie-com/togo/internal/model"
	"github.com/manabie-com/togo/internal/storages/ent"
	"github.com/manabie-com/togo/internal/storages/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func TestAuthService_Login(t *testing.T) {
	mockUserRepos := &mocks.UserRepositoryMock{}
	authService := NewAuthService("", mockUserRepos)
	ctx := context.Background()
	req := model.LoginCredential{
		UserName: "demo",
		Password: "example",
	}

	mockUserRepos.On("FindByUsername", ctx, req.UserName).Return(&ent.User{Password: "$2a$10$DZjcIPh9cv.cWH62dYII0uaYsPjvSCR4hMfMBNl4GrSaktw7vaQ2O"}, nil)

	foundUser, err := authService.Login(ctx, req)

	assert.Nil(t, err)
	assert.NotNil(t, foundUser)

	mockUserRepos.AssertExpectations(t)
}

func TestAuthService_CreateUser(t *testing.T) {
	mockUserRepos := &mocks.UserRepositoryMock{}
	authService := NewAuthService("", mockUserRepos)
	ctx := context.Background()

	mockUserRepos.On("CreateUser", mock.Anything, mock.Anything, mock.Anything).Return(&ent.User{}, nil)

	newUser, err := authService.CreateUser(ctx, "ad", "pwd")

	assert.Nil(t, err)
	assert.NotNil(t, newUser)

	mockUserRepos.AssertExpectations(t)
}
