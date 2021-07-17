package helpers

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHashPassword(t *testing.T) {
	assert := assert.New(t)
	password, err := HashPassword("123456")

	assert.NotEmpty(password)
	assert.NoError(err)
}

func TestCheckPasswordHash(t *testing.T) {
	assert := assert.New(t)
	password, _ := HashPassword("123456")

	assert.True(CheckPasswordHash("123456", password))
	assert.False(CheckPasswordHash("1234456", password))
}
