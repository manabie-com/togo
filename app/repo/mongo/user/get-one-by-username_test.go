package user_test

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestGetUserByUsername(t *testing.T) {
	user1 := CreateRandomUser(t)
	user2, err := userRepoInstance.GetOneByUsername(context.Background(), user1.Username)
	require.NoError(t, err)
	require.NotEmpty(t, user2)

	// compare user1 == user2 ?
	require.Equal(t, user1.Username, user2.Username)
	require.Equal(t, user1.HashedPassword, user2.HashedPassword)
	require.Equal(t, user1.MaxTasks, user2.MaxTasks)
	require.Equal(t, user1.CurrentTasks, user2.CurrentTasks)
	require.WithinDuration(t, *user1.CreatedDate, *user2.CreatedDate, time.Second)
}
