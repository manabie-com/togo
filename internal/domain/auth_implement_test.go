package domain_test

import (
	"context"
	"errors"
	"github.com/manabie-com/togo/internal/domain"
	"github.com/manabie-com/togo/internal/domain/entity"
	"github.com/manabie-com/togo/mocks"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

var sampleToken = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6ImNhMTEyZWE0LTQxODItMTFlYy1hYWU0LWU0NTRlODI1MmY1MiIsInVzZXJuYW1lIjoibXR1YW4iLCJpc3N1ZWRBdCI6IjIwMjEtMTEtMTBUMDA6MzA6NTMuMzMyMzI0ODQxKzA3OjAwIiwiZXhwaXJlZEF0IjoiMjAyMS0xMS0xMFQwMDozMTo1My4zMzIzMjQ5MjMrMDc6MDAifQ.gDPuf3oaxA8UkTwfpDdG9nGwS1JSQGg_LZKApYctdNc"

func TestSignInSuccess(t *testing.T) {
	username := "mtuan"
	mockAuthRepository := new(mocks.UserRepository)
	mockTokenMarker := new(mocks.Token)
	tokenConfig := domain.NewTokenConfig(time.Minute * 15)
	mockAuthRepository.On("GetUser", mock.Anything, username).
		Return(&entity.User{
			Username:       username,
			HashedPassword: "$2a$10$3jEtynoYdZJlw2fTUMjuCeGxHEjvc8a23gXMaidDW3yKPjMWFbb4W",
		}, nil)
	mockTokenMarker.On("CreateToken", username, time.Minute*15).
		Return(sampleToken, nil)
	uc := domain.NewAuthUseCase(mockAuthRepository, mockTokenMarker, tokenConfig)
	accessToken, err := uc.SignIn(context.TODO(), username, "example")
	require.NoError(t, err)
	require.NotNil(t, accessToken)
}

func TestSignInGetUserFailed(t *testing.T) {
	username := "mtuan"
	mockAuthRepository := new(mocks.UserRepository)
	mockTokenMarker := new(mocks.Token)
	tokenConfig := domain.NewTokenConfig(time.Minute * 15)
	mockAuthRepository.On("GetUser", mock.Anything, username).
		Return(nil, domain.UserNotFound)
	mockTokenMarker.On("CreateToken", username, time.Minute*15).
		Return(sampleToken, nil)
	uc := domain.NewAuthUseCase(mockAuthRepository, mockTokenMarker, tokenConfig)
	accessToken, err := uc.SignIn(context.TODO(), username, "example")
	require.Error(t, err)
	require.EqualError(t, err, domain.UserNotFound.Error())
	require.Equal(t, "", accessToken)
}

func TestSignInCreateTokenFailed(t *testing.T) {
	username := "mtuan"
	mockAuthRepository := new(mocks.UserRepository)
	mockTokenMarker := new(mocks.Token)
	tokenConfig := domain.NewTokenConfig(time.Minute * 15)
	mockAuthRepository.On("GetUser", mock.Anything, username).
		Return(&entity.User{
			Username:       username,
			HashedPassword: "$2a$10$3jEtynoYdZJlw2fTUMjuCeGxHEjvc8a23gXMaidDW3yKPjMWFbb4W",
		}, nil)
	mockTokenMarker.On("CreateToken", username, time.Minute*15).
		Return("", errors.New("create token failed"))
	uc := domain.NewAuthUseCase(mockAuthRepository, mockTokenMarker, tokenConfig)
	accessToken, err := uc.SignIn(context.TODO(), username, "example")
	require.Error(t, err)
	require.Equal(t, "", accessToken)
}

func TestVerifyTokenError(t *testing.T) {
	username := "mtuan"
	mockAuthRepository := new(mocks.UserRepository)
	mockTokenMarker := new(mocks.Token)
	tokenConfig := domain.NewTokenConfig(time.Minute * 15)
	mockAuthRepository.On("GetUser", mock.Anything, username).
		Return(&entity.User{
			Username:       username,
			HashedPassword: "blabla",
		}, nil)
	mockTokenMarker.On("CreateToken", username, time.Minute*15).
		Return(sampleToken, nil)
	uc := domain.NewAuthUseCase(mockAuthRepository, mockTokenMarker, tokenConfig)
	accessToken, err := uc.SignIn(context.TODO(), username, "example")
	require.Error(t, err)
	require.Equal(t, "", accessToken)
	require.EqualError(t, err, domain.WrongPassword.Error())

}
