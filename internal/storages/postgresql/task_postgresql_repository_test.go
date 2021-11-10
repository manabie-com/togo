package postgresql_test

import (
	"context"
	"errors"
	"github.com/manabie-com/togo/internal/storages/postgresql"
	"github.com/stretchr/testify/require"
	"gopkg.in/DATA-DOG/go-sqlmock.v1"
	"regexp"
	"testing"
)

func TestCountTaskInDayByUsernameSuccess(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	username := "firstUser"
	columns := []string{"count"}
	mock.ExpectPrepare(regexp.QuoteMeta("SELECT count(*) FROM tasks WHERE username = $1 AND DATE(created_at) = DATE(NOW());")).ExpectQuery().
		WithArgs(username).
		WillReturnRows(sqlmock.NewRows(columns).AddRow(5))
	repo := postgresql.NewTaskPostgresqlRepository(db)
	count, err := repo.CountTaskInDayByUsername(context.TODO(), username)
	require.NoError(t, err)
	require.Equal(t, 5, count)
}

func TestCountTaskInDayByUsernameReturnError(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	username := "firstUser"
	mock.ExpectPrepare(regexp.QuoteMeta("SELECT count(*) FROM tasks WHERE username = $1 AND DATE(created_at) = DATE(NOW());")).ExpectQuery().
		WithArgs(username).
		WillReturnError(errors.New("something was wrong"))
	repo := postgresql.NewTaskPostgresqlRepository(db)
	count, err := repo.CountTaskInDayByUsername(context.TODO(), username)
	require.Error(t, err)
	require.Equal(t, 0, count)
}

func TestCountTaskInDayUserUserErrorPrepare(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	username := "firstUser"
	mock.ExpectPrepare(regexp.QuoteMeta("SELECT count(*) FROM tasks WHERE username = $1 AND DATE(created_at) = DATE(NOW());")).
		WillReturnError(errors.New("something was wrong"))
	repo := postgresql.NewTaskPostgresqlRepository(db)
	count, err := repo.CountTaskInDayByUsername(context.TODO(), username)
	require.Error(t, err)
	require.Equal(t, 0, count)
}

func TestCreateTaskSuccess(t *testing.T) {
	db, mock, err := sqlmock.New()
	content := "content"
	username := "firstUser"
	require.NoError(t, err)
	mock.ExpectPrepare(regexp.QuoteMeta("INSERT INTO tasks(content, username) VALUES ($1, $2)")).ExpectExec().
		WithArgs(content, username).
		WillReturnResult(sqlmock.NewResult(1, 1))
	repo := postgresql.NewTaskPostgresqlRepository(db)
	err = repo.Create(context.TODO(), content, username)
	require.NoError(t, err)
}

func TestCreateTaskReturnError(t *testing.T) {
	db, mock, err := sqlmock.New()
	content := "content"
	username := "firstUser"
	require.NoError(t, err)
	mock.ExpectPrepare(regexp.QuoteMeta("INSERT INTO tasks(content, username) VALUES ($1, $2)")).ExpectExec().
		WithArgs(content, username).
		WillReturnError(errors.New("something was wrong"))
	repo := postgresql.NewTaskPostgresqlRepository(db)
	err = repo.Create(context.TODO(), content, username)
	require.Error(t, err)
}

func TestCreateTaskReturnErrorPrepare(t *testing.T) {
	db, mock, err := sqlmock.New()
	content := "content"
	username := "firstUser"
	require.NoError(t, err)
	mock.ExpectPrepare(regexp.QuoteMeta("INSERT INTO tasks(content, username) VALUES ($1, $2)")).
		WillReturnError(errors.New("something was wrong"))
	repo := postgresql.NewTaskPostgresqlRepository(db)
	err = repo.Create(context.TODO(), content, username)
	require.Error(t, err)
}
