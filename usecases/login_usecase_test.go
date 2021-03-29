package usecases

import (
	"context"
	"errors"
	"github.com/manabie-com/togo/domains"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLoginUseCaseSuccess(t *testing.T) {
	ctx := context.Background()
	expected := "AAAAA"

	mockRepo := new(DBMock)
	mockRepo.On("VerifyUser", ctx, &domains.LoginRequest{}).Return(&domains.User{}, nil)
	mockAuth := new(AuthMock)
	mockAuth.On("CreateToken", int64(0)).Return(expected, nil)

	uc := NewLoginUseCase(mockRepo, mockAuth)
	tokenResult, err := uc.Execute(ctx, &LoginInput{})

	mockAuth.AssertExpectations(t)
	mockRepo.AssertExpectations(t)
	assert.Equal(t, expected, tokenResult)
	assert.Nil(t, err)
}

func TestLoginUseCaseErrorInvalidUsernameOrPassword(t *testing.T) {
	ctx := context.Background()
	expected := ""

	mockRepo := new(DBMock)
	mockRepo.On("VerifyUser", ctx, &domains.LoginRequest{}).Return(nil, domains.ErrorNotFound)
	mockAuth := new(AuthMock)

	uc := NewLoginUseCase(mockRepo, mockAuth)
	tokenResult, err := uc.Execute(ctx, &LoginInput{})

	mockAuth.AssertExpectations(t)
	mockRepo.AssertExpectations(t)
	assert.Equal(t, expected, tokenResult)
	assert.Equal(t, ErrorInvalidUsernameOrPassword, err)
}

func TestLoginUseCaseErrorInternalError(t *testing.T) {
	ctx := context.Background()
	expected := errors.New("internal server error")

	mockRepo := new(DBMock)
	mockRepo.On("VerifyUser", ctx, &domains.LoginRequest{}).Return(nil, expected)
	mockAuth := new(AuthMock)

	uc := NewLoginUseCase(mockRepo, mockAuth)
	tokenResult, err := uc.Execute(ctx, &LoginInput{})

	mockAuth.AssertExpectations(t)
	mockRepo.AssertExpectations(t)
	assert.Equal(t, expected, err)
	assert.Equal(t, "", tokenResult)
}

func TestLoginUseCaseErrorCreateTokenError(t *testing.T) {
	ctx := context.Background()
	expected := errors.New("create token error")

	mockRepo := new(DBMock)
	mockRepo.On("VerifyUser", ctx, &domains.LoginRequest{}).Return(&domains.User{}, nil)
	mockAuth := new(AuthMock)
	mockAuth.On("CreateToken", int64(0)).Return("", expected)

	uc := NewLoginUseCase(mockRepo, mockAuth)
	tokenResult, err := uc.Execute(ctx, &LoginInput{})

	mockAuth.AssertExpectations(t)
	mockRepo.AssertExpectations(t)
	assert.Equal(t, expected, err)
	assert.Equal(t, "", tokenResult)
}
