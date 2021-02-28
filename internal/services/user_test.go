package services

import (
	"context"
	"testing"
	"github.com/stretchr/testify/assert"
)

const JWTKey = "123456"

func TestUserIDFromCtx(t *testing.T)  {
	s := &UserService{
		JWTKey: JWTKey,
	}
	ctx := context.WithValue(context.Background(), UserAuthKey(1), "firstUser")
	userId, ok := s.UserIDFromCtx(ctx)

	assert := assert.New(t)
	assert.Equal(false, ok)
	assert.Empty(userId)
}

func TestUserIDFromCtxError(t *testing.T)  {
	s := &UserService{
		JWTKey: JWTKey,
	}
	ctx := context.WithValue(context.Background(), UserAuthKey(0), "firstUser")
	userId, ok := s.UserIDFromCtx(ctx)

	assert := assert.New(t)
	assert.Equal(true, ok)
	assert.Equal("firstUser", userId)
}

func TestValidToken(t *testing.T) {
	s := &UserService{
		JWTKey: JWTKey,
	}
	token, _ := s.createToken("firstUser")
	validToken, err := s.ValidToken(token)

	assert := assert.New(t)
	assert.NoError(err)
	assert.NotEqual(validToken, token)
}

func TestInvalidToken(t *testing.T) {
	s := &UserService{
		JWTKey: JWTKey,
	}
	token, _ := s.createToken("firstUser")
	invalidToken, err := s.ValidToken("123")

	assert := assert.New(t)
	assert.Error(err)
	assert.NotEqual(invalidToken, token)
}

func TestGetAuthToken(t *testing.T) {
	storage := new(storeMock)
	s := &UserService{
		JWTKey: JWTKey,
		Storage: storage,
	}
	ctx := context.Background()

	storage.On("ValidateUser", ctx, value("firstUser"), value("example")).Return(true)
	token, err := s.GetAuthToken(ctx, "firstUser", "example")

	assert := assert.New(t)
	assert.NoError(err)
	assert.NotEmpty(token)
}

func TestGetAuthTokenError(t *testing.T) {
	storage := new(storeMock)
	s := &UserService{
		JWTKey: JWTKey,
		Storage: storage,
	}
	ctx := context.Background()

	storage.On("ValidateUser", ctx, value("firstUser"), value("example")).Return(false)
	token, err := s.GetAuthToken(ctx, "firstUser", "example")

	assert := assert.New(t)
	assert.Error(err)
	assert.Empty(token)
}

func TestCreateToken(t *testing.T) {
	s := &UserService{
		JWTKey: JWTKey,
	}
	token, err := s.createToken("firstUser")

	assert := assert.New(t)
	assert.NoError(err)
	assert.NotEmpty(token)
}

