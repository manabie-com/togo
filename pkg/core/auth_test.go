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
	token := `eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2MTY5MDA2ODcsInN1YiI6MTExMTF9.PEFk_EYX13NJAQTVzyhs99tw91VzFmeV4hRvk96gRm0`
	os.Setenv("JWT_TOKEN",  "A JWT TOKEN")

	auth, _ := NewAppAuthenticator()
	userId, err := auth.ValidateToken(&http.Request{Header: map[string][]string{
		"Authorization": []string{token},
	}})
	assert.NotEmpty(t, userId)
	assert.Nil(t, err)
}
