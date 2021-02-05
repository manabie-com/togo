package services

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_userIDFromCtx(t *testing.T) {
	ctx := context.WithValue(context.Background(), userAuthKey(0), "firstUser")
	ID, ok := userIDFromCtx(ctx)

	assert.Equal(t, ok, true)
	assert.Equal(t, ID, "firstUser")
}

func Test_userIDFromCtx_Error(t *testing.T) {
	ctx := context.WithValue(context.Background(), userAuthKey(1), "firstUser")
	ID, ok := userIDFromCtx(ctx)

	assert.Equal(t, ok, false)
	assert.Equal(t, ID, "")
}

func Test_createToken(t *testing.T) {
	s := ToDoService{}
	token, err := s.createToken("firstUser")
	assert.NoError(t, err)
	assert.NotEqual(t, token, "")
}

func Test_validToken(t *testing.T) {
	s := ToDoService{}
	token, err := s.createToken("firstUser")
	assert.NoError(t, err)
	assert.NotNil(t, token)

	request := httptest.NewRequest(http.MethodGet, "/login?user_id=firstUser", nil)
	request.Header.Set("Authorization", token)

	_, ok := s.validToken(request)
	assert.Equal(t, ok, true)
}

func Test_validToken_Error(t *testing.T) {
	s := ToDoService{}
	request := httptest.NewRequest(http.MethodGet, "/login?user_id=firstUser", nil)
	request.Header.Set("Authorization", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2MTE5MzAxMzIsInVzZXJfaWQiOiJmaXJzdFVzZXIifQ.QxYb0gYbu3a25GCADZEQxMfmHqsUIilfCOApaC1GNhc")

	_, ok := s.validToken(request)
	assert.Equal(t, ok, false)
}
