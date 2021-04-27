package postgres_test

import (
	"context"
	"fmt"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/manabie-com/togo/internal/app/models"
	"github.com/manabie-com/togo/internal/storages/postgres"
	"github.com/stretchr/testify/assert"
)

func TestValidateUserSuccess(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	rows := sqlmock.
		NewRows([]string{"id", "username", "password"}).
		AddRow(100, "ngohuuphong", "phong123")
	mock.
		ExpectQuery("SELECT id, username, password FROM users WHERE username = \\$1 AND password = \\$2").
		WithArgs("ngohuuphong", "phong123").
		WillReturnRows(rows)

	u := postgres.NewPosgresql(db)
	user, err := u.ValidateUser(context.Background(), "ngohuuphong", "phong123")
	assert.NoError(t, err)
	assert.Equal(t, user.ID, uint64(100))
}

func TestValidateUserFailed(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	mock.ExpectQuery("SELECT id, username, password FROM users WHERE username = \\$1 AND password = \\$2").
		WithArgs("user123", "password123").
		WillReturnError(fmt.Errorf("some error"))

	u := postgres.NewPosgresql(db)
	_, err = u.ValidateUser(context.Background(), "user123", "password123")
	assert.NotNil(t, err)
}

func TestRetrieveTasksSuccess(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	mockTask := []*models.Task{
		{
			ID:          111,
			Content:     "statement 1",
			UserID:      100,
			CreatedDate: "2021-04-26",
		},
		{
			ID:          222,
			Content:     "statement 2",
			UserID:      100,
			CreatedDate: "2021-04-26",
		},
		{
			ID:          333,
			Content:     "statement 3",
			UserID:      100,
			CreatedDate: "2021-04-26",
		},
	}

	rows := sqlmock.
		NewRows([]string{"id", "content", "user_id", "created_date"}).
		AddRow(111, "statement 1", 100, "2021-04-26").
		AddRow(222, "statement 2", 100, "2021-04-26").
		AddRow(333, "statement 3", 100, "2021-04-26")
	mock.
		ExpectQuery("SELECT id, content, user_id, created_date FROM tasks WHERE user_id = \\$1 AND created_date::date = \\$2::date").
		WithArgs(100, "2021-04-26").
		WillReturnRows(rows)

	u := postgres.NewPosgresql(db)
	tasks, err := u.RetrieveTasks(context.Background(), uint64(100), "2021-04-26")
	assert.NoError(t, err)
	assert.Equal(t, tasks, mockTask)
}

func TestRetrieveTasksFailedWrongUserID(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	mockTask := []*models.Task{
		{
			ID:          111,
			Content:     "statement 1",
			UserID:      100,
			CreatedDate: "2021-04-26",
		},
		{
			ID:          222,
			Content:     "statement 2",
			UserID:      100,
			CreatedDate: "2021-04-26",
		},
		{
			ID:          333,
			Content:     "statement 3",
			UserID:      100,
			CreatedDate: "2021-04-26",
		},
	}

	rows := sqlmock.
		NewRows([]string{"id", "content", "user_id", "created_date"}).
		AddRow(111, "statement 1", 100, "2021-04-26").
		AddRow(222, "statement 2", 100, "2021-04-26").
		AddRow(333, "statement 3", 100, "2021-04-26")
	mock.
		ExpectQuery("SELECT id, content, user_id, created_date FROM tasks WHERE user_id = \\$1 AND created_date::date = \\$2::date").
		WithArgs(100, "2021-04-26").
		WillReturnRows(rows)

	u := postgres.NewPosgresql(db)
	tasks, _ := u.RetrieveTasks(context.Background(), 200, "2021-04-26")
	assert.NotEqual(t, tasks, mockTask)
}

func TestRetrieveTasksFailedWrongUserCreateDate(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	mockTask := []*models.Task{
		{
			ID:          111,
			Content:     "statement 1",
			UserID:      100,
			CreatedDate: "2021-04-26",
		},
		{
			ID:          222,
			Content:     "statement 2",
			UserID:      100,
			CreatedDate: "2021-04-26",
		},
		{
			ID:          333,
			Content:     "statement 3",
			UserID:      100,
			CreatedDate: "2021-04-26",
		},
	}

	rows := sqlmock.
		NewRows([]string{"id", "content", "user_id", "created_date"}).
		AddRow(111, "statement 1", 100, "2021-04-26").
		AddRow(222, "statement 2", 100, "2021-04-26").
		AddRow(333, "statement 3", 100, "2021-04-26")
	mock.
		ExpectQuery("SELECT id, content, user_id, created_date FROM tasks WHERE user_id = \\$1 AND created_date::date = \\$2::date").
		WithArgs(100, "2021-04-26").
		WillReturnRows(rows)

	u := postgres.NewPosgresql(db)
	tasks, _ := u.RetrieveTasks(context.Background(), 100, "2021-09-09")
	assert.NotEqual(t, tasks, mockTask)
}

func TestAddTaskSuccess(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	mockTask := &models.Task{
		ID:          100,
		Content:     "statement added",
		UserID:      200,
		CreatedDate: "2021-04-26",
	}

	query := regexp.QuoteMeta(`INSERT INTO tasks (id, content, user_id, created_date) SELECT $1, $2, $3, $4::date FROM users u WHERE u.id = $3 AND (SELECT COUNT(id) FROM tasks WHERE user_id = $3 AND created_date::date = $4::date) < u.max_todo RETURNING id`)
	prep := mock.ExpectPrepare(query)
	prep.ExpectExec().WithArgs(mockTask.ID, mockTask.Content, mockTask.UserID, mockTask.CreatedDate).WillReturnResult(sqlmock.NewResult(1, 1))

	u := postgres.NewPosgresql(db)
	err = u.AddTask(context.Background(), mockTask)
	assert.NoError(t, err)
}

func TestAddTaskFailed(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	mockTask := &models.Task{
		ID:          100,
		Content:     "statement added",
		UserID:      200,
		CreatedDate: "2021-04-26",
	}
	rowsAffected := 0

	query := regexp.QuoteMeta(`INSERT INTO tasks (id, content, user_id, created_date) SELECT $1, $2, $3, $4::date FROM users u WHERE u.id = $3 AND (SELECT COUNT(id) FROM tasks WHERE user_id = $3 AND created_date::date = $4::date) < u.max_todo RETURNING id`)
	prep := mock.ExpectPrepare(query)
	prep.ExpectExec().WithArgs(mockTask.ID, mockTask.Content, mockTask.UserID, mockTask.CreatedDate).WillReturnResult(sqlmock.NewResult(0, int64(rowsAffected)))
	var actualErr string
	if rowsAffected < 1 {
		actualErr = "the task daily limit is reached"
	}
	u := postgres.NewPosgresql(db)
	err = u.AddTask(context.Background(), mockTask)
	assert.EqualError(t, err, actualErr)
}
