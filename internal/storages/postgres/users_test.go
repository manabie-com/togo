package postgres

import (
	"context"
	"testing"

	"github.com/jericogantuangco/togo/internal/util"
	"github.com/stretchr/testify/require"
)

func createRandomPassword(number int) string {
	return util.RandomString(number)
}

func TestCreateUser(t *testing.T) {
	lenOfPassword := 5
	lenOfString := 5
	arg := CreateUserParams{
		Username: util.RandomString(lenOfString),
		Password: createRandomPassword(lenOfPassword),
	}

	user, err := testQueries.CreateUser(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, user)
	require.Len(t, user.Password, lenOfPassword)
}
