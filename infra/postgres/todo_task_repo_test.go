// +build integration

package postgres_test

import (
	"context"
	"testing"
	"time"

	"github.com/laghodessa/togo/domain"
	"github.com/laghodessa/togo/domain/todo"
	"github.com/laghodessa/togo/infra/postgres"
	"github.com/laghodessa/togo/test/todofixture"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestTodoTaskRepo_AddTask(t *testing.T) {
	migrate(t)
	t.Cleanup(clearDB)

	repo := postgres.NewTodoTaskRepo(db)

	loc, err := time.LoadLocation("Asia/Ho_Chi_Minh")
	require.NoError(t, err)

	err = repo.AddTask(context.Background(), todofixture.NewTask(), loc, 1)
	assert.ErrorContains(t, err, `null value in column "user_pk"`)

	err = repo.AddTask(context.Background(), todofixture.NewTask(), loc, 0)
	assert.ErrorIs(t, err, todo.ErrUserHitTaskDailyLimit)
}

func TestTodoTaskRepo_CountInTimeRangeByUserID(t *testing.T) {
	migrate(t)
	t.Cleanup(clearDB)

	repo := postgres.NewTodoTaskRepo(db)
	count, err := repo.CountInTimeRangeByUserID(context.Background(), domain.NewID(), time.Now(), time.Now())
	require.NoError(t, err)
	assert.Equal(t, 0, count)
}
