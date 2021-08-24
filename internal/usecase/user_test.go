package usecase

import (
	"context"
	"errors"
	"testing"

	adapter_mocks "github.com/manabie-com/togo/internal/adapter/mocks"
	"github.com/manabie-com/togo/internal/common"
	"github.com/manabie-com/togo/internal/dto"
	sqlite_mocks "github.com/manabie-com/togo/internal/storages/sqlite/mocks"
	"github.com/stretchr/testify/assert"
)

var (
	loginReqDTO = &dto.LoginRequestDTO{
		UserID:   "firstUser",
		Password: "example",
	}
	pwdMd5 = "1a79a4d60de6718e8e5b326e338ae533"
)

func TestLoginSuccess(t *testing.T) {
	var (
		wantResp = &dto.LoginResponseDTO{
			Token: "token",
		}
		liteDB = new(sqlite_mocks.LiteDB)
		jwt    = new(adapter_mocks.JWTAdapter)
	)
	liteDB.On("ValidateUser", context.Background(), loginReqDTO.UserID, pwdMd5).Return(true)
	jwt.On("CreateToken", context.Background(), loginReqDTO.UserID).Return("token", nil)

	userUsecase := NewUserUsecase(jwt, liteDB)
	resp, _ := userUsecase.Login(context.Background(), loginReqDTO)
	if !assert.Equal(t, wantResp, resp) {
		t.Errorf("userUsecase.Login resp = %v, want = %v", resp, wantResp)
	}
}

func TestLoginUserIDPasswordEmptyError(t *testing.T) {
	var (
		req = &dto.LoginRequestDTO{
			UserID:   "",
			Password: "",
		}
		wantErr = errors.New(common.ReasonUserIDPasswordEmptyError.Code())
		liteDB  = new(sqlite_mocks.LiteDB)
		jwt     = new(adapter_mocks.JWTAdapter)
	)

	userUsecase := NewUserUsecase(jwt, liteDB)
	_, err := userUsecase.Login(context.Background(), req)
	if !assert.Equal(t, wantErr, err) {
		t.Errorf("userUsecase.Login err = %v, want = %v", err, wantErr)
	}
}

func TestLoginUnauthorizedError(t *testing.T) {
	var (
		wantErr = errors.New(common.ReasonUnauthorized.Code())
		liteDB  = new(sqlite_mocks.LiteDB)
		jwt     = new(adapter_mocks.JWTAdapter)
	)
	liteDB.On("ValidateUser", context.Background(), loginReqDTO.UserID, pwdMd5).Return(false)

	userUsecase := NewUserUsecase(jwt, liteDB)
	_, err := userUsecase.Login(context.Background(), loginReqDTO)
	if !assert.Equal(t, wantErr, err) {
		t.Errorf("userUsecase.Login err = %v, want = %v", err, wantErr)
	}
}

func TestLoginReasonCreateTokenError(t *testing.T) {
	var (
		wantErr = errors.New(common.ReasonCreateTokenError.Code())
		liteDB  = new(sqlite_mocks.LiteDB)
		jwt     = new(adapter_mocks.JWTAdapter)
	)
	liteDB.On("ValidateUser", context.Background(), loginReqDTO.UserID, pwdMd5).Return(true)
	jwt.On("CreateToken", context.Background(), loginReqDTO.UserID).Return("", errors.New(common.ReasonCreateTokenError.Code()))

	userUsecase := NewUserUsecase(jwt, liteDB)
	_, err := userUsecase.Login(context.Background(), loginReqDTO)
	if !assert.Equal(t, wantErr, err) {
		t.Errorf("userUsecase.Login err = %v, want = %v", err, wantErr)
	}
}

func TestVerifyTokenSuccess(t *testing.T) {
	var (
		req = &dto.VerifyTokenRequestDTO{
			Token: "Bearer token",
		}
		wantResp = &dto.VerifyTokenResponseDTO{
			UserID: "firstUser",
		}
		liteDB = new(sqlite_mocks.LiteDB)
		jwt    = new(adapter_mocks.JWTAdapter)
	)
	jwt.On("VerifyToken", context.Background(), "token").Return("firstUser", nil)

	userUsecase := NewUserUsecase(jwt, liteDB)
	resp, _ := userUsecase.VerifyToken(context.Background(), req)
	if !assert.Equal(t, wantResp, resp) {
		t.Errorf("userUsecase.VerifyToken resp = %v, want = %v", resp, wantResp)
	}
}

func TestVerifyTokenNotMatchPattern(t *testing.T) {
	var (
		req = &dto.VerifyTokenRequestDTO{
			Token: "token",
		}
		wantErr = errors.New(common.ReasonInvalidToken.Code())
		liteDB  = new(sqlite_mocks.LiteDB)
		jwt     = new(adapter_mocks.JWTAdapter)
	)

	userUsecase := NewUserUsecase(jwt, liteDB)
	_, err := userUsecase.VerifyToken(context.Background(), req)
	if !assert.Equal(t, wantErr, err) {
		t.Errorf("userUsecase.VerifyToken err = %v, want = %v", err, wantErr)
	}
}

func TestVerifyTokenEmpty(t *testing.T) {
	var (
		req = &dto.VerifyTokenRequestDTO{
			Token: "Bearer ",
		}
		wantErr = errors.New(common.ReasonInvalidToken.Code())
		liteDB  = new(sqlite_mocks.LiteDB)
		jwt     = new(adapter_mocks.JWTAdapter)
	)

	userUsecase := NewUserUsecase(jwt, liteDB)
	_, err := userUsecase.VerifyToken(context.Background(), req)
	if !assert.Equal(t, wantErr, err) {
		t.Errorf("userUsecase.VerifyToken err = %v, want = %v", err, wantErr)
	}
}

func TestVerifyTokenInvalidToken(t *testing.T) {
	var (
		req = &dto.VerifyTokenRequestDTO{
			Token: "Bearer token",
		}
		wantErr = errors.New(common.ReasonInvalidToken.Code())
		liteDB  = new(sqlite_mocks.LiteDB)
		jwt     = new(adapter_mocks.JWTAdapter)
	)
	jwt.On("VerifyToken", context.Background(), "token").Return("", errors.New(common.ReasonInvalidToken.Code()))

	userUsecase := NewUserUsecase(jwt, liteDB)
	_, err := userUsecase.VerifyToken(context.Background(), req)
	if !assert.Equal(t, wantErr, err) {
		t.Errorf("userUsecase.VerifyToken err = %v, want = %v", err, wantErr)
	}
}
