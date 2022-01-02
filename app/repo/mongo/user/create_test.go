package user_test

import (
	"context"
	"testing"

	"github.com/manabie-com/togo/app/model"
	"github.com/manabie-com/togo/app/utils"

	userRepo "github.com/manabie-com/togo/app/repo/mongo/user"

	"github.com/stretchr/testify/require"
)

func CreateRandomUser(t *testing.T) model.User {
	// random info request
	hashedPassword, err := utils.HashPassword(utils.RandomString(6))
	require.NoError(t, err)

	arg := userRepo.CreateReq{
		Username:       utils.RandomOwner(),
		HashedPassword: hashedPassword,
		MaxTasks:       int(utils.RandomInt(1, 10)),
		// tracing
		CreatedIP: "127.0.0.1",
	}

	user, err := userRepoInstance.Create(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, user)

	require.Equal(t, arg.Username, user.Username)
	require.Equal(t, arg.HashedPassword, user.HashedPassword)
	require.Equal(t, arg.MaxTasks, user.MaxTasks)
	require.Zero(t, user.CurrentTasks)
	require.True(t, user.ChangedPasswordAt == nil)
	require.NotZero(t, user.CreatedDate)

	return user
}

func TestCreateUser(t *testing.T) {
	CreateRandomUser(t)
}
