package usecase_test

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
			MaxTodo:      100,
		}
		_, err := th.Usecase.Store.User().Create(context.Background(), &u)
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
					ID:      uuid.New().String(),
					Content: fake.Sentence(),
					UserID:  uuid.New().String(),
				},
				true,
			},
		}

		for _, tc := range tcs {
			actual, err := th.Usecase.AddTask(context.Background(), tc.task.UserID, &tc.task)
			if tc.hasError {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
				assert.NotNil(t, actual.CreatedDate)
			}
		}
	})

	t.Run("When add task limited", func(t *testing.T) {
		u := model.User{
			ID:           uuid.New().String(),
			PasswordHash: uuid.New().String(),
			MaxTodo:      3,
		}
		_, err := th.Usecase.Store.User().Create(context.Background(), &u)
		require.NoError(t, err)

		// add tasks
		for i := 0; i < u.MaxTodo; i++ {
			task := model.Task{
				ID:      uuid.New().String(),
				Content: fake.Sentence(),
				UserID:  u.ID,
			}
			_, err := th.Usecase.AddTask(context.Background(), task.UserID, &task)
			require.NoError(t, err)
		}

		// add failed
		task := model.Task{
			ID:      uuid.New().String(),
			Content: fake.Sentence(),
			UserID:  u.ID,
		}
		_, err = th.Usecase.AddTask(context.Background(), task.UserID, &task)
		require.Error(t, err)
	})

	t.Run("When list tasks", func(t *testing.T) {
		u1 := model.User{
			ID:           uuid.New().String(),
			PasswordHash: uuid.New().String(),
			MaxTodo:      6,
		}
		_, err := th.Usecase.Store.User().Create(context.Background(), &u1)
		require.NoError(t, err)

		tasks := map[string]*model.Task{
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

		// add tasks
		for k, task := range tasks {
			task.ID = k
			actual, err := th.Usecase.AddTask(context.Background(), task.UserID, task)
			require.NoError(t, err)
			tasks[k].CreatedDate = actual.CreatedDate
		}

		// compare result to input
		actual, err := th.Usecase.ListTasks(
			context.Background(),
			u1.ID,
			sql.NullString{
				String: time.Now().UTC().Format("2006-01-02"),
				Valid:  true,
			},
		)
		require.NoError(t, err)
		assert.Len(t, actual, len(tasks))

		// get all tasks
		actual, err = th.Usecase.ListTasks(
			context.Background(),
			u1.ID,
			sql.NullString{},
		)
		require.NoError(t, err)
		assert.Len(t, actual, len(tasks))
	})
}
