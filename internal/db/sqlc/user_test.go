package db

import (
	"context"
	"github.com/stretchr/testify/require"
	"task-manage/internal/utils"
	"testing"
	"time"
)

func createRandomUser(t *testing.T) (User, error) {
	hashedPassword, err := utils.HashPassword(utils.RandomString(6))
	require.NoError(t, err)

	arg := CreateUserParams{
		UserName:         utils.RandomString(4),
		HashedPassword:   hashedPassword,
		CreatedAt:        time.Now(),
		UpdatedAt:        time.Now(),
		MaximumTaskInDay: 1,
	}

	user, err := testQueries.CreateUser(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, user)

	require.Equal(t, arg.UserName, user.UserName)
	require.Equal(t, arg.HashedPassword, user.HashedPassword)
	return user, err
}

func TestCreateUser(t *testing.T) {
	createRandomUser(t)
}
