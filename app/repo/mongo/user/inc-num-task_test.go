package user_test

import (
	"context"
	"testing"

	userRepo "github.com/manabie-com/togo/app/repo/mongo/user"
	"github.com/stretchr/testify/require"
)

func TestIncNumTask(t *testing.T) {

	testCase := []struct {
		name            string
		overloadMaxTask bool
	}{
		{
			name:            "inc_to_max",
			overloadMaxTask: false,
		},
		{
			name:            "inc_to_overload_max",
			overloadMaxTask: true,
		},
	}

	for _, tc := range testCase {
		t.Run(tc.name, func(t *testing.T) {
			user := createRandomUser(t)

			incTaskReq := userRepo.IncNumTaskReq{
				UserID:   user.ID,
				MaxTasks: user.MaxTasks,
			}

			currentTask := 0
			for i := 0; i < user.MaxTasks; i++ {
				currentTask++

				// inc user's task
				updatedUser, err := userRepoInstance.IncNumTask(context.Background(), incTaskReq)
				require.NoError(t, err)
				require.Equal(t, updatedUser.CurrentTasks, currentTask)
				require.True(t, updatedUser.CurrentTasks <= user.MaxTasks)
			}

			if tc.overloadMaxTask {
				updatedUser, err := userRepoInstance.IncNumTask(context.Background(), incTaskReq)
				require.NoError(t, err)
				// Empty result updated user, because not adpapt condition query currentTask < maxTask
				require.Empty(t, updatedUser)
			}
		})
	}
}
