package core

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"os"
	"testing"
)

func TestAppAuthenticate_CreateToken(t *testing.T) {
	os.Setenv("JWT_TOKEN",  "A JWT TOKEN")
	auth, _ := NewAppAuthenticator()
	token, _ := auth.CreateToken(11111)
	assert.NotEmpty(t, token)
}

func TestAppAuthenticate_ValidateToken(t *testing.T) {
	token := `eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MDMzODcxMDMsInN1YiI6MTExMTF9.0zMbkrg3rXZbVUEUYkOJBATZpwwnn7dCmfM9nDHUKRo`
	os.Setenv("JWT_TOKEN",  "A JWT TOKEN")

	auth, _ := NewAppAuthenticator()
	userId, err := auth.ValidateToken(&http.Request{Header: map[string][]string{
		"Authorization": []string{token},
	}})
	assert.NotEmpty(t, userId)
	assert.Nil(t, err)
}
