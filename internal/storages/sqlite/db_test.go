package sqllite

import (
	"context"
	"database/sql"
	"log"
	"testing"
	"togo/internal/storages"

	_ "github.com/moemoe89/go-unit-test-sql/repository"
	"github.com/stretchr/testify/assert"

	"github.com/DATA-DOG/go-sqlmock"
)

var task = &storages.Task{
	ID:          "e1da0b9b-7ecc-44f9-82ff-4623cc50446a",
	Content:     "first content",
	UserID:      "firstUser",
	CreatedDate: "2020-06-29",
}

var userID = sql.NullString{
	String: "firstUser",
	Valid:  true,
}

var pwd = sql.NullString{
	String: "example",
	Valid:  true,
}

var createdDate = sql.NullString{
	String: "2020-06-29",
	Valid:  true,
}

func NewMock() (*sql.DB, sqlmock.Sqlmock) {
	db, mock, err := sqlmock.New()
	if err != nil {
		log.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	return db, mock
}

func Test_RetrieveTasks(t *testing.T) {
	db, mock := NewMock()

	l := LiteDB{
		DB: db,
	}

	defer func() {
		l.DB.Close()
	}()

	stmt := "SELECT id, content, user_id, created_date FROM tasks WHERE user_id = \\? AND created_date = \\?"

	rows := sqlmock.NewRows([]string{"id", "content", "user_id", "created_date"}).AddRow(task.ID, task.Content, task.UserID, task.CreatedDate)
	mock.ExpectQuery(stmt).WillReturnRows(rows)

	ctx := context.Background()
	tasks, err := l.RetrieveTasks(ctx, userID, createdDate)
	assert.NotEmpty(t, tasks)
	assert.NoError(t, err)
	assert.Len(t, tasks, 1)
}

func Test_AddTask(t *testing.T) {
	db, mock := NewMock()

	l := LiteDB{
		DB: db,
	}

	defer func() {
		l.DB.Close()
	}()

	stmt := "INSERT INTO tasks \\(id, content, user_id, created_date\\) VALUES \\(\\?, \\?, \\?, \\?\\)"
	mock.ExpectExec(stmt).WithArgs(task.ID, task.Content, task.UserID, task.CreatedDate).WillReturnResult(sqlmock.NewResult(0, 1))

	ctx := context.Background()
	err := l.AddTask(ctx, task)
	assert.NoError(t, err)
}

func Test_AddTask_Error(t *testing.T) {
	db, mock := NewMock()

	l := LiteDB{
		DB: db,
	}

	defer func() {
		l.DB.Close()
	}()

	stmt := "INSERT INTO tasks \\(id, content, user_id, created_date\\) VALUES \\(\\?, \\?, \\?, \\?\\)"
	prep := mock.ExpectPrepare(stmt)
	prep.ExpectExec().WithArgs(task.ID, task.Content, task.UserID, task.CreatedDate).WillReturnResult(sqlmock.NewResult(0, 0))

	ctx := context.Background()
	err := l.AddTask(ctx, task)
	assert.Error(t, err)
}

func Test_ValidateUser(t *testing.T) {
	db, mock := NewMock()

	l := LiteDB{
		DB: db,
	}

	defer func() {
		l.DB.Close()
	}()

	stmt := "SELECT id FROM users WHERE id = \\? AND password = \\?"

	rows := sqlmock.NewRows([]string{"id"}).AddRow(task.UserID)
	mock.ExpectQuery(stmt).WithArgs(userID, pwd).WillReturnRows(rows)

	ctx := context.Background()
	ok := l.ValidateUser(ctx, userID, pwd)
	assert.Equal(t, ok, true)
}

func Test_ValidateUser_Error(t *testing.T) {
	db, mock := NewMock()

	l := LiteDB{
		DB: db,
	}

	defer func() {
		l.DB.Close()
	}()

	stmt := "SELECT id FROM users WHERE id = \\? AND password = \\?"

	rows := sqlmock.NewRows([]string{"id"})
	mock.ExpectQuery(stmt).WithArgs(userID, pwd).WillReturnRows(rows)

	ctx := context.Background()
	ok := l.ValidateUser(ctx, userID, pwd)
	assert.Equal(t, ok, false)
}
