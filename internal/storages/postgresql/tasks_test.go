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
	"time"
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

		const totalTasks = 2
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

		// count total tasks
		count, err := th.Store.Task().CountTasksByUser(context.Background(), u.ID, sql.NullString{})
		require.NoError(t, err)
		assert.EqualValues(t, totalTasks, count)
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

		tasksOfU1 := map[string]*model.Task{
			uuid.New().String(): {
				Content: fake.Sentence(),
				UserID:  u1.ID,
			},
			uuid.New().String(): {
				Content: fake.Sentence(),
				UserID:  u1.ID,
			},
			uuid.New().String(): {
				Content: fake.Sentence(),
				UserID:  u1.ID,
			},
		}

		tasksOfU2 := map[string]*model.Task{
			uuid.New().String(): {
				Content: fake.Sentence(),
				UserID:  u2.ID,
			},
			uuid.New().String(): {
				Content: fake.Sentence(),
				UserID:  u2.ID,
			},
		}

		// add tasks to db
		for k, task := range tasksOfU1 {
			task.ID = k
			actual, err := th.Store.Task().AddTask(context.Background(), task.UserID, task)
			require.NoError(t, err)
			tasksOfU1[k].CreatedDate = actual.CreatedDate
		}

		// add tasks to db
		for k, task := range tasksOfU2 {
			task.ID = k
			actual, err := th.Store.Task().AddTask(context.Background(), task.UserID, task)
			require.NoError(t, err)
			tasksOfU2[k].CreatedDate = actual.CreatedDate
		}

		// compare result to input
		actual1, err := th.Store.Task().RetrieveTasks(
			context.Background(),
			u1.ID,
			sql.NullString{
				String: time.Now().UTC().Format("2006-01-02"),
				Valid:  true,
			},
		)
		require.NoError(t, err)
		assert.Len(t, actual1, len(tasksOfU1))
		for _, act := range actual1 {
			expected, ok := tasksOfU1[act.ID]
			assert.True(t, ok)
			assert.EqualValues(t, expected, act)
		}
		// get all tasks
		actual1, err = th.Store.Task().RetrieveTasks(
			context.Background(),
			u1.ID,
			sql.NullString{},
		)
		require.NoError(t, err)
		assert.Len(t, actual1, len(tasksOfU1))

		actual2, err := th.Store.Task().RetrieveTasks(
			context.Background(),
			u2.ID,
			sql.NullString{
				String: time.Now().UTC().Format("2006-01-02"),
				Valid:  true,
			},
		)
		require.NoError(t, err)
		assert.Len(t, actual2, len(tasksOfU2))
		for _, act := range actual2 {
			expected, ok := tasksOfU2[act.ID]
			assert.True(t, ok)
			assert.EqualValues(t, expected, act)
		}
		// get all tasks
		actual2, err = th.Store.Task().RetrieveTasks(
			context.Background(),
			u2.ID,
			sql.NullString{},
		)
		require.NoError(t, err)
		assert.Len(t, actual2, len(tasksOfU2))
	})
}
