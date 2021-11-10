package postgresql_test

import (
	"context"
	"database/sql"
	"errors"
	"github.com/manabie-com/togo/internal/domain"
	"github.com/manabie-com/togo/internal/storages/postgresql"
	"github.com/stretchr/testify/require"
	"gopkg.in/DATA-DOG/go-sqlmock.v1"
	"regexp"
	"testing"
)

func TestGetUserSuccess(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	username := "firstUser"
	columns := []string{"username", "hashed_password", "max_todo"}
	mock.ExpectPrepare(regexp.QuoteMeta("SELECT username, hashed_password, max_todo FROM users WHERE username = $1")).
		ExpectQuery().WithArgs(username).WillReturnRows(sqlmock.NewRows(columns).AddRow(username, "123456", 5))
	authRepo := postgresql.NewUserPostgresqlRepository(db)
	user, err := authRepo.GetUser(context.TODO(), username)
	require.NoError(t, err)
	require.NotNil(t, user)
	require.Equal(t, username, user.Username)
	require.Equal(t, "123456", user.HashedPassword)
}

func TestGetUserErrorPrepare(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	username := "firstUser"
	mock.ExpectPrepare(regexp.QuoteMeta("SELECT username, hashed_password, max_todo FROM users WHERE username = $1")).
		WillReturnError(errors.New("something was wrong"))
	authRepo := postgresql.NewUserPostgresqlRepository(db)
	user, err := authRepo.GetUser(context.TODO(), username)
	require.Error(t, err)
	require.Nil(t, user)
}

func TestGetUserErrorQueryRow(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	username := "firstUser"
	mock.ExpectPrepare(regexp.QuoteMeta("SELECT username, hashed_password, max_todo FROM users WHERE username = $1")).ExpectQuery().WithArgs(username).
		WillReturnError(errors.New("something was wrong"))
	authRepo := postgresql.NewUserPostgresqlRepository(db)
	user, err := authRepo.GetUser(context.TODO(), username)
	require.Error(t, err)
	require.Nil(t, user)
}

func TestGetUserReturnErrNoRows(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	username := "firstUser"
	mock.ExpectPrepare(regexp.QuoteMeta("SELECT username, hashed_password, max_todo FROM users WHERE username = $1")).ExpectQuery().WithArgs(username).
		WillReturnError(sql.ErrNoRows)
	authRepo := postgresql.NewUserPostgresqlRepository(db)
	user, err := authRepo.GetUser(context.TODO(), username)
	require.Error(t, err)
	require.EqualError(t, err, domain.UserNotFound.Error())
	require.Nil(t, user)
}
