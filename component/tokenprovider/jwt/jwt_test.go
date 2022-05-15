package jwt_test

import (
	"github.com/japananh/togo/component/tokenprovider"
	"github.com/japananh/togo/component/tokenprovider/jwt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestJwtProvider_Generate(t *testing.T) {
	secretKey := "secretKey"
	payload := tokenprovider.TokenPayload{UserId: 1}
	expiry := 86400 // 24 hours in milliseconds
	token, err := jwt.NewTokenJWTProvider(secretKey).Generate(payload, expiry)

	require.Nil(t, err, err)
	require.NotNil(t, token.Token)
	require.NotNil(t, token.Created)
	assert.Equal(t, token.Expiry, expiry, "they should be equal")
}

func TestJwtProvider_Validate(t *testing.T) {
	secretKey := "secretKey"
	jwtProvider := jwt.NewTokenJWTProvider(secretKey)
	userId := 1
	tkPayload := tokenprovider.TokenPayload{UserId: userId}
	expiry := 86400 // 24 hours in milliseconds

	token, err := jwtProvider.Generate(tkPayload, expiry)
	require.Nil(t, err, err)

	payload, err := jwtProvider.Validate(token.Token)
	require.Nil(t, err, err)
	assert.Equal(t, payload.UserId, userId, "they should be equal")
}
