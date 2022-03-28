// +build integration

package postgres_test

import (
	"context"
	"testing"

	"github.com/laghodessa/togo/domain"
	"github.com/laghodessa/togo/infra/postgres"
	"github.com/laghodessa/togo/test/todofixture"
	"github.com/stretchr/testify/assert"
)

func TestTodoUserRepo_GetUser(t *testing.T) {
	migrate(t)
	t.Cleanup(clearDB)

	repo := postgres.NewTodoUserRepo(db)

	_, err := repo.GetUser(context.Background(), domain.NewID())
	assert.ErrorIs(t, err, domain.ErrNotFound)
}

func TestTodoUserRepo_AddUser(t *testing.T) {
	migrate(t)
	t.Cleanup(clearDB)

	repo := postgres.NewTodoUserRepo(db)

	err := repo.AddUser(context.Background(), todofixture.NewUser())
	assert.NoError(t, err)
}
