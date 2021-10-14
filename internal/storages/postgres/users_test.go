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
	user, err := testQueries.CreateUser(context.Background(), createRandomPassword(lenOfPassword))
	require.NoError(t, err)
	require.NotEmpty(t,user)
	require.Len(t, user.Password, lenOfPassword)
}
