package postgresql_test

import (
	"context"
	"database/sql"
	"github.com/google/uuid"
	"github.com/icrowley/fake"
	"github.com/manabie-com/togo/internal/model"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestTasks(t *testing.T) {
	th, err := Setup()
	require.NoError(t, err)
	defer th.Teardown()

	t.Run("When add task", func(t *testing.T) {
		u := model.User{
			ID:           uuid.New().String(),
			PasswordHash: uuid.New().String(),
			MaxTodo:      6,
		}
		_, err := th.Store.User().Create(context.Background(), &u)
		require.NoError(t, err)

		tcs := []struct {
			task     model.Task
			hasError bool
		}{
			{
				model.Task{
					ID:      uuid.New().String(),
					Content: fake.Sentence(),
					UserID:  u.ID,
				},
				false,
			},
			{
				model.Task{
					ID:      uuid.New().String(),
					Content: fake.Sentence(),
					UserID:  u.ID,
				},
				false,
			},
			{
				model.Task{
					ID:      uuid.New().String(),
					Content: fake.Sentence(),
				},
				true,
			},
			{
				model.Task{
					Content: fake.Sentence(),
					UserID:  u.ID,
				},
				true,
			},
			{
				model.Task{
					ID:      uuid.New().String(),
					Content: fake.Sentence(),
					UserID:  uuid.New().String(),
				},
				true,
			},
		}

		for _, tc := range tcs {
			actual, err := th.Store.Task().AddTask(context.Background(), tc.task.UserID, &tc.task)
			if tc.hasError {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
				assert.NotNil(t, actual.CreatedDate)
			}
		}
	})

	t.Run("When retrieve tasks", func(t *testing.T) {
		u1 := model.User{
			ID:           uuid.New().String(),
			PasswordHash: uuid.New().String(),
			MaxTodo:      6,
		}
		_, err := th.Store.User().Create(context.Background(), &u1)
		require.NoError(t, err)

		u2 := model.User{
			ID:           uuid.New().String(),
			PasswordHash: uuid.New().String(),
			MaxTodo:      3,
		}
		_, err = th.Store.User().Create(context.Background(), &u2)
		require.NoError(t, err)

		tasksOfU1 := []*model.Task{
			{
				ID:      uuid.New().String(),
				Content: fake.Sentence(),
				UserID:  u1.ID,
			},
			{
				ID:      uuid.New().String(),
				Content: fake.Sentence(),
				UserID:  u1.ID,
			},
			{
				ID:      uuid.New().String(),
				Content: fake.Sentence(),
				UserID:  u1.ID,
			},
		}

		tasksOfU2 := []*model.Task{
			{
				ID:      uuid.New().String(),
				Content: fake.Sentence(),
				UserID:  u2.ID,
			},
			{
				ID:      uuid.New().String(),
				Content: fake.Sentence(),
				UserID:  u2.ID,
			},
			{
				ID:      uuid.New().String(),
				Content: fake.Sentence(),
				UserID:  u2.ID,
			},
		}

		// add tasks to db
		for i, task := range tasksOfU1 {
			actual, err := th.Store.Task().AddTask(context.Background(), task.UserID, task)
			require.NoError(t, err)
			tasksOfU1[i].CreatedDate = actual.CreatedDate
		}

		// add tasks to db
		for i, task := range tasksOfU2 {
			actual, err := th.Store.Task().AddTask(context.Background(), task.UserID, task)
			require.NoError(t, err)
			tasksOfU2[i].CreatedDate = actual.CreatedDate
		}

		th.Store.Task().RetrieveTasks(context.Background(), sql.NullString{
			String: u1.ID,
			Valid:  true,
		}, )
	})
}
