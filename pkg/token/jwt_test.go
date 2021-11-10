package token_test

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/manabie-com/togo/pkg/token"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestJWTMaker(t *testing.T) {
	maker := token.NewJWTMaker("12345")
	username := "mtuan"
	duration := time.Minute

	issuedAt := time.Now()
	expiredAt := time.Now().Add(duration)

	jwtToken, err := maker.CreateToken(username, duration)
	require.NoError(t, err)
	require.NotEmpty(t, jwtToken)

	payload, err := maker.VerifyToken(jwtToken)
	require.NoError(t, err)
	require.NotEmpty(t, payload)

	require.NotZero(t, payload.Id)
	require.Equal(t, username, payload.Username)
	require.WithinDuration(t, issuedAt, payload.IssuedAt, time.Second)
	require.WithinDuration(t, expiredAt, payload.ExpiredAt, time.Second)
}

func TestExpiredJWTToken(t *testing.T) {
	maker := token.NewJWTMaker("12345")
	username := "mtuan"
	duration := -time.Minute
	jwtToken, err := maker.CreateToken(username, duration)
	require.NoError(t, err)
	require.NotEmpty(t, jwtToken)

	payload, err := maker.VerifyToken(jwtToken)
	require.EqualError(t, err, token.ErrorExpiredToken.Error())
	require.Nil(t, payload)
}

func TestInvalidJWTToken(t *testing.T) {
	payload, err := token.NewPayload("mtuan", time.Minute)
	require.NoError(t, err)

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodNone, payload)
	tokenStr, err := jwtToken.SignedString(jwt.UnsafeAllowNoneSignatureType)
	require.NoError(t, err)
	maker := token.NewJWTMaker("12345")

	payload, err = maker.VerifyToken(tokenStr)
	require.Error(t, err)
	require.EqualError(t, err, token.ErrorInvalidToken.Error())
	require.Nil(t, payload)
}
