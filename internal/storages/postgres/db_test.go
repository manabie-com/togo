package postgres

import (
	"context"
	"database/sql"
	"errors"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/manabie-com/togo/internal/storages"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestPostgreDB_RetrieveTasks_Happy(t *testing.T) {
	id := "id"
	content := "content"
	userID := "user_id"
	createdDate := "created_date"

	db, mock, err := sqlmock.New()
	require.Nil(t, err)
	defer db.Close()

	mock.ExpectQuery(
		`SELECT id, content, user_id, created_date FROM tasks WHERE user_id = .* AND created_date = .*`).
		WillReturnRows(sqlmock.NewRows([]string{"id", "content", "user_id", "created_date"}).
			AddRow(id, content, userID, createdDate).
			AddRow(id, content, userID, createdDate))

	ctx := context.Background()
	postgreDB := &PostgreDB{DB: db}

	tasks, err := postgreDB.RetrieveTasks(ctx,
		sql.NullString{
			String: userID,
			Valid:  true,
		}, sql.NullString{
			String: createdDate,
			Valid:  true,
		})

	require.Nil(t, err)
	require.Equal(t, []*storages.Task{{
		ID:          id,
		Content:     content,
		UserID:      userID,
		CreatedDate: createdDate,
	}, {
		ID:          id,
		Content:     content,
		UserID:      userID,
		CreatedDate: createdDate,
	}}, tasks)
}

func TestPostgreDB_RetrieveTasks_Empty(t *testing.T) {
	userID := "user_id"
	createdDate := "created_date"

	db, mock, err := sqlmock.New()
	require.Nil(t, err)
	defer db.Close()

	mock.ExpectQuery(
		`SELECT id, content, user_id, created_date FROM tasks WHERE user_id = .* AND created_date = .*`).
		WillReturnRows(sqlmock.NewRows([]string{"id", "content", "user_id", "created_date"}))

	ctx := context.Background()
	postgreDB := &PostgreDB{DB: db}

	_, err = postgreDB.RetrieveTasks(ctx,
		sql.NullString{
			String: userID,
			Valid:  true,
		}, sql.NullString{
			String: createdDate,
			Valid:  true,
		})

	require.Nil(t, err)
}

func TestPostgreDB_RetrieveTasks_Edge(t *testing.T) {
	userID := "user_id"
	createdDate := "created_date"

	db, mock, err := sqlmock.New()
	require.Nil(t, err)
	defer db.Close()

	mock.ExpectQuery(
		`SELECT id, content, user_id, created_date FROM tasks WHERE user_id = .* AND created_date = .*`).
		WillReturnError(errors.New("random error"))

	ctx := context.Background()
	postgreDB := &PostgreDB{DB: db}

	_, err = postgreDB.RetrieveTasks(ctx,
		sql.NullString{
			String: userID,
			Valid:  true,
		}, sql.NullString{
			String: createdDate,
			Valid:  true,
		})

	require.NotNil(t, err)
}

func TestPostgreDB_AddTask_Happy(t *testing.T) {
	id := "id"
	content := "content"
	userID := "user_id"
	createdDate := "created_date"

	db, mock, err := sqlmock.New()
	require.Nil(t, err)
	defer db.Close()

	mock.ExpectExec(
		`INSERT INTO tasks`).
		WithArgs(id, content, userID, createdDate).
		WillReturnResult(sqlmock.NewResult(1, 1))

	ctx := context.Background()
	postgreDB := &PostgreDB{DB: db}

	_, err = postgreDB.AddTask(ctx,
		&storages.Task{
			ID:          id,
			Content:     content,
			UserID:      userID,
			CreatedDate: createdDate,
		})

	require.Nil(t, err)
}

func TestPostgreDB_AddTask_Edge(t *testing.T) {
	id := "id"
	content := "content"
	userID := "user_id"
	createdDate := "created_date"

	db, mock, err := sqlmock.New()
	require.Nil(t, err)
	defer db.Close()

	mock.ExpectExec(
		`INSERT INTO tasks`).
		WithArgs(id, content, userID, createdDate).WillReturnError(errors.New("random_error"))

	ctx := context.Background()
	postgreDB := &PostgreDB{DB: db}

	_, err = postgreDB.AddTask(ctx,
		&storages.Task{
			ID:          id,
			Content:     content,
			UserID:      userID,
			CreatedDate: createdDate,
		})

	require.NotNil(t, err)
}

func TestPostgreDB_ValidateUser_Happy(t *testing.T) {
	userID := "user_id"
	password := "password"

	db, mock, err := sqlmock.New()
	require.Nil(t, err)
	defer db.Close()

	mock.ExpectQuery(
		`SELECT id FROM users WHERE id = .* AND password = .*`).
		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow("inputted_id"))

	ctx := context.Background()
	postgreDB := &PostgreDB{DB: db}

	valid := postgreDB.ValidateUser(ctx,
		sql.NullString{
			String: userID,
			Valid:  true,
		}, sql.NullString{
			String: password,
			Valid:  true,
		})
	require.True(t, valid)
}

func TestPostgreDB_ValidateUser_Edge(t *testing.T) {
	userID := "user_id"
	password := "password"

	db, mock, err := sqlmock.New()
	require.Nil(t, err)
	defer db.Close()

	mock.ExpectQuery(
		`SELECT id FROM users WHERE id = .* AND password = .*`).
		WillReturnRows(sqlmock.NewRows([]string{"id"}))

	ctx := context.Background()
	postgreDB := &PostgreDB{DB: db}

	valid := postgreDB.ValidateUser(ctx,
		sql.NullString{
			String: userID,
			Valid:  true,
		}, sql.NullString{
			String: password,
			Valid:  true,
		})
	require.False(t, valid)
}
