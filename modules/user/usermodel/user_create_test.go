package usermodel_test

import (
	"github.com/japananh/togo/modules/user/usermodel"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestUserCreate_Validate(t *testing.T) {
	user := usermodel.UserCreate{}
	user.Email = "user@gmail.com"
	user.Password = "user@123"

	err := user.Validate()
	require.Nil(t, err, err)
}

func TestUserCreate_VerifyPassword(t *testing.T) {
	var tsc = []struct {
		arg      string
		expected string
	}{
		{"password@123", ""},
		{"pass", usermodel.ErrNotEnoughCharacterMsg},
		{"password 1234", usermodel.ErrInvalidCharacterMsg},
		{"12345678", usermodel.ErrMustHaveLetterMsg},
		{"password", usermodel.ErrMustHaveNumberMsg},
		{"password123", usermodel.ErrMustHaveSpecialCharacterMsg},
	}
	for _, tc := range tsc {
		output := usermodel.VerifyPassword(tc.arg)
		assert.Equal(t, output, tc.expected, "they should be equal")
	}
}
