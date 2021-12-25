package utils

import (
	"testing"

	"github.com/bxcodec/faker/v3"
	"github.com/stretchr/testify/assert"
)

func Test_Password(t *testing.T) {
	password := faker.Password()
	hashedPassword, err := HashPassword(password)

	assert.NoError(t, err)
	assert.NotEmpty(t, hashedPassword)

	err = VerifyPassword(password, hashedPassword)
	assert.NoError(t, err)
}

func Test_WrongPassword(t *testing.T) {
	password := faker.Password()
	hashedPassword, err := HashPassword(password)

	assert.NoError(t, err)
	assert.NotEmpty(t, hashedPassword)

	otherPassword := faker.Password()
	err = VerifyPassword(otherPassword, hashedPassword)
	assert.Error(t, err)
}
