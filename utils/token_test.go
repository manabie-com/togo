package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_CreateAndCheckToken(t *testing.T) {
	id := "someID"
	jwtKey := "someKey"

	token, err := CreateToken(id, jwtKey)
	assert.NoError(t, err)
	assert.NotEmpty(t, token)

	result, ok := ValidateToken(token, jwtKey)
	assert.True(t, ok)
	assert.Equal(t, id, result)
}

func Test_WrongToken(t *testing.T) {
	jwtKey := "someKey"

	token := "somethingtoken"

	id, ok := ValidateToken(token, jwtKey)
	assert.False(t, ok)
	assert.Empty(t, id)
}
