package util

import (
	"testing"

	"github.com/stretchr/testify/require"
	"golang.org/x/crypto/bcrypt"
)

func TestPassword(t *testing.T) {
	password := RandomPassword()
	hashedPassword1, err := HashPassword(password, bcrypt.DefaultCost)
	require.NoError(t, err)
	require.NotEmpty(t, hashedPassword1)
	err = CheckPassword(password, hashedPassword1)
	require.NoError(t, err)
	wrongPassword := RandomPassword()
	err = CheckPassword(wrongPassword, hashedPassword1)
	require.EqualError(t, err, bcrypt.ErrMismatchedHashAndPassword.Error())
	hashedPassword2, err := HashPassword(password, bcrypt.DefaultCost)
	require.NoError(t, err)
	require.NotEmpty(t, hashedPassword2)
	require.NotEqual(t, hashedPassword1, hashedPassword2)
	hashedPassword3, err := HashPassword(password, bcrypt.MaxCost+1)
	require.Error(t, err)
	require.Empty(t, hashedPassword3)
}
