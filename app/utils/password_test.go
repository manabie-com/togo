package utils_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/manabie-com/togo/app/utils"
)

func TestPasswords(t *testing.T) {
	testCases := []struct {
		name             string
		rawPassword      string
		wrongPassword    string
		errCheckPassword bool
	}{
		{
			name:             "hashed_check_true_password",
			rawPassword:      utils.RandomString(6),
			wrongPassword:    utils.RandomString(6),
			errCheckPassword: false,
		},
		{
			name:             "hashed_check_wrong_password",
			rawPassword:      utils.RandomString(6),
			wrongPassword:    utils.RandomString(6),
			errCheckPassword: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			hashedPassword, err := utils.HashPassword(tc.rawPassword)
			require.NoError(t, err)
			require.NotEmpty(t, hashedPassword)

			if !tc.errCheckPassword {
				err = utils.CheckPassword(tc.rawPassword, hashedPassword)
				require.NoError(t, err)
			} else {
				err = utils.CheckPassword(tc.wrongPassword, hashedPassword)
				require.Error(t, err)
			}
		})
	}
}
