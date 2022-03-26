package todo_test

import (
	"testing"

	"github.com/laghodessa/togo/domain/todo"
	"github.com/laghodessa/togo/test/todofixture"
	"github.com/stretchr/testify/require"
)

func TestUser_HitTaskDailyLimit(t *testing.T) {
	user := todofixture.NewUser(func(u *todo.User) {
		u.TaskDailyLimit = 3
	})

	require.Nil(t, user.HitTaskDailyLimit(2))
	require.ErrorIs(t, user.HitTaskDailyLimit(3), todo.ErrUserHitTaskDailyLimit)
	require.ErrorIs(t, user.HitTaskDailyLimit(4), todo.ErrUserHitTaskDailyLimit)
}
