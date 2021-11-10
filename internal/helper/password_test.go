package helper_test

import (
	"github.com/manabie-com/togo/internal/helper"
	"github.com/stretchr/testify/require"
	"golang.org/x/crypto/bcrypt"
	"testing"
)

func TestPasswordHelper(t *testing.T) {
	password := "example"
	hashedPassword, err := helper.HashedPassword(password)
	require.NoError(t, err)
	require.NotEmpty(t, hashedPassword)
	err = helper.CheckPassword(password, hashedPassword)
	require.NoError(t, err)

	wrongPassword := "example1"
	err = helper.CheckPassword(wrongPassword, hashedPassword)
	require.Error(t, err)
	require.EqualError(t, err, bcrypt.ErrMismatchedHashAndPassword.Error())
}
