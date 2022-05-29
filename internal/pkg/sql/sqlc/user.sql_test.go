package db

import (
	"context"
	"testing"

	"github.com/dinhquockhanh/togo/internal/pkg/util"
	"github.com/stretchr/testify/require"
)

func createRandomUser(t *testing.T) *User {
	// TODO: create tier, limit
	hashedPwd, err := util.HashPassword(util.RandomString(10))
	require.NoError(t, err)

	arg := CreateUserParams{
		Username:       util.RandomOwner(),
		HashedPassword: hashedPwd,
		FullName:       util.RandomOwner(),
		Email:          util.RandomEmail(),
		TierID:         int32(util.RandomInt(1, 2)),
	}

	user, err := testQueries.CreateUser(context.Background(), &arg)
	require.NoError(t, err)
	require.NotEmpty(t, user)

	require.NotEmpty(t, user.Username)
	require.NotEmpty(t, user.HashedPassword)
	require.NotEmpty(t, user.FullName)
	require.NotEmpty(t, user.Email)

	return user
}

func TestQueries_CreateUserIntegration(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test")
	}
	createRandomUser(t)
}

func TestQueries_GetUserIntegration(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test")
	}
	user := createRandomUser(t)
	res, err := testQueries.GetUserByName(context.Background(), user.Username)
	require.NoError(t, err)
	require.Equal(t, user.Username, res.Username)
	require.Equal(t, user.HashedPassword, res.HashedPassword)
	require.Equal(t, user.FullName, res.FullName)
	require.Equal(t, user.Email, res.Email)
	require.NotZero(t, user.CreatedAt)
}
