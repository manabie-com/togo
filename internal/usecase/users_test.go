package usecase_test

import (
	"context"
	"github.com/google/uuid"
	"github.com/manabie-com/togo/internal/model"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestUsers(t *testing.T) {
	th, err := Setup()
	require.NoError(t, err)
	defer th.Teardown()

	t.Run("When get a user", func(t *testing.T) {
		u := model.User{
			ID:           uuid.New().String(),
			PasswordHash: uuid.New().String(),
			MaxTodo:      6,
		}
		_, err := th.Usecase.Store.User().Create(context.Background(), &u)
		require.NoError(t, err)

		actual, err := th.Usecase.GetUser(context.Background(), u.ID)
		require.NoError(t, err)
		assert.EqualValues(t, u, *actual)

		actual, err = th.Usecase.GetUser(context.Background(), uuid.New().String())
		require.Error(t, err)
	})
}
