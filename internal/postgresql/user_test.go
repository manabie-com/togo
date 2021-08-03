package postgresql

import (
	"context"
	"testing"
	"time"

	"github.com/bxcodec/faker/v3"
	"github.com/stretchr/testify/require"
)

func CreateRandomUser(t *testing.T) User {
	arg := InsertUserParams{
		Username: faker.Username(),
		Password: faker.Password(),
	}

	user, err := testQueries.InsertUser(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, user)

	require.Equal(t, arg.Username, user.Username)
	require.Equal(t, arg.Password, user.Password)
	require.NotZero(t, user.CreatedAt)

	return user
}

func TestCreateUser(t *testing.T) {
	CreateRandomUser(t)
}

func TestQueries_GetUserByUsername(t *testing.T) {
	user1 := CreateRandomUser(t)
	user2, err := testQueries.GetUserByUsername(context.Background(), user1.Username)
	require.NoError(t, err)
	require.NotEmpty(t, user2)

	require.Equal(t, user1.Username, user2.Username)
	require.Equal(t, user1.Password, user2.Password)
	require.WithinDuration(t, user1.CreatedAt, user2.CreatedAt, time.Second)
}

func TestQueries_GetUser(t *testing.T) {
	user1 := CreateRandomUser(t)
	user2, err := testQueries.GetUser(context.Background(), user1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, user2)

	require.Equal(t, user1.Username, user2.Username)
	require.Equal(t, user1.Password, user2.Password)
	require.WithinDuration(t, user1.CreatedAt, user2.CreatedAt, time.Second)
}
