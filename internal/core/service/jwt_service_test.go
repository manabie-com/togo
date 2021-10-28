package service

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateTokenSuccess(t *testing.T) {
	expectedUserId := "user_id_example"

	svc := NewJwtService("wqGyEBBfPK9w3Lxw")
	token, err := svc.CreateToken(expectedUserId)
	assert.NoError(t, err, "should not be error")
	assert.NotEmpty(t, token, "should not be empty")

	userId, err := svc.ParseToken(token)
	assert.NoError(t, err, "should not be error")
	assert.Equal(t, expectedUserId, userId, "parse token failed")
}

func TestCreateTokenFailed(t *testing.T) {
	expectedUserId := "user_id_example"

	jwtCreateToken := NewJwtService("wqGyEBBfPK9w3Lxw")
	token, err := jwtCreateToken.CreateToken(expectedUserId)
	assert.NoError(t, err, "should not be error")
	assert.NotEmpty(t, token, "should not be empty")

	jwtParseToken := NewJwtService("wrong_key")
	userId, err := jwtParseToken.ParseToken(token)
	assert.Error(t, err, "should be error")
	assert.NotEqual(t, expectedUserId, userId, "should be failed")
}
