package db

import (
	"context"
	"testing"
	"togo/util"

	"github.com/stretchr/testify/require"
	"golang.org/x/crypto/bcrypt"
)

func createRandomUser(t *testing.T) User {
	hashedPassword, err := util.HashPassword(util.RandomPassword(), bcrypt.DefaultCost)
	require.NoError(t, err)
	arg := CreateUserParams{
		Username:       util.RandomName(),
		HashedPassword: hashedPassword,
		FullName:       util.RandomName(),
		Email:          util.RandomEmail(),
	}
	user, err := testQueries.CreateUser(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, user)
	require.Equal(t, arg.Username, user.Username)
	require.Equal(t, arg.HashedPassword, user.HashedPassword)
	require.Equal(t, arg.FullName, user.FullName)
	require.Equal(t, arg.Email, user.Email)
	require.True(t, user.PasswordChangeAt.IsZero())
	require.NotZero(t, user.CreatedAt)
	return user
}

func updateRandomedUser(t *testing.T, username string, dailyCap int64, dailyQuantity int64) {
	updatedUser, err := testQueries.UpdateUserDailyQuantity(context.Background(), UpdateUserDailyQuantityParams{
		Username:      username,
		DailyQuantity: dailyQuantity,
	})
	require.NoError(t, err)
	require.NotEmpty(t, updatedUser)
	require.Equal(t, updatedUser.DailyQuantity, dailyQuantity)
	updatedUser, err = testQueries.UpdateUserDailyCap(context.Background(), UpdateUserDailyCapParams{
		Username: username,
		DailyCap: dailyCap,
	})
	require.NoError(t, err)
	require.NotEmpty(t, updatedUser)
	require.Equal(t, updatedUser.DailyCap, dailyCap)
}
