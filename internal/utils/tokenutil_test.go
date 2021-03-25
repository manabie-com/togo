package utils

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCreateToken(t *testing.T) {
	jwtKey := "aGFwcHkyZ290b2dvCg=="
	userId := "firstUser"

	token, err := CreateToken(userId, jwtKey)

	require.Nil(t, err)

	require.NotNil(t, token)
}
