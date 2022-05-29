package util

import (
	"testing"

	"github.com/stretchr/testify/require"
	"golang.org/x/crypto/bcrypt"
)

func TestHashPassword(t *testing.T) {
	pwd := RandomString(15)
	hashedPwd, err := HashPassword(pwd)
	require.NoError(t, err)
	require.NotEmpty(t, hashedPwd)

	err = CheckPassword(pwd, hashedPwd)
	require.NoError(t, err)

	wrongPwd := RandomString(15)
	err = CheckPassword(wrongPwd, hashedPwd)
	require.EqualError(t, err, bcrypt.ErrMismatchedHashAndPassword.Error())

	hashedPwd2, err := HashPassword(pwd)
	require.NoError(t, err)
	require.NotEmpty(t, hashedPwd2)
	require.NotEqual(t, hashedPwd, hashedPwd2)
}
