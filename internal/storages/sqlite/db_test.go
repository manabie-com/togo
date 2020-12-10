package sqllite

import (
	"context"
	"database/sql"
	"log"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/manabie-com/togo/internal/storages"
	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/assert"
)

var (
	testUser        = storages.User{ID: "userid", Password: "password"}
	testCreatedDate = "2020-01-01"
	testTaskList    = []*storages.Task{
		{ID: "task-id-1", Content: "task-content-1", UserID: testUser.ID, CreatedDate: testCreatedDate},
		{ID: "task-id-2", Content: "task-content-2", UserID: testUser.ID, CreatedDate: testCreatedDate},
	}
)

func newSQLMock() (*LiteDB, sqlmock.Sqlmock) {
	db, mock, err := sqlmock.New()
	if err != nil {
		log.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	return &LiteDB{DB: db}, mock
}

func ns(s string) sql.NullString {
	return sql.NullString{String: s, Valid: true}
}

func TestValidateUserSuccess(t *testing.T) {
	ldb, mock := newSQLMock()
	defer ldb.DB.Close()

	ctx := context.Background()
	query := `SELECT id FROM users WHERE id = (.+) AND password = (.+)`
	rows := sqlmock.NewRows([]string{"id"}).AddRow(testUser.ID)
	mock.ExpectQuery(query).WithArgs(testUser.ID, testUser.Password).WillReturnRows(rows)
	assert.True(t, ldb.ValidateUser(ctx, ns(testUser.ID), ns(testUser.Password)))
}

func TestValidateUserFailure(t *testing.T) {
	ldb, mock := newSQLMock()
	defer ldb.DB.Close()

	ctx := context.Background()
	query := `SELECT id FROM users WHERE id = (.+) AND password = (.+)`
	rows := sqlmock.NewRows([]string{"id"})
	mock.ExpectQuery(query).WillReturnRows(rows)
	assert.False(t, ldb.ValidateUser(ctx, ns("wrongUserID"), ns(testUser.Password)))
	assert.False(t, ldb.ValidateUser(ctx, ns(testUser.ID), ns("wrongPassword")))
}

func TestRetreiveTaskSuccess(t *testing.T) {
	ldb, mock := newSQLMock()
	defer ldb.DB.Close()

	ctx := context.Background()
	query := `SELECT id, content, user_id, created_date FROM tasks WHERE user_id = (.+) AND created_date = (.+)`
	rows := sqlmock.NewRows([]string{"id", "content", "user_id", "created_date"})
	for _, v := range testTaskList {
		rows.AddRow(v.ID, v.Content, v.UserID, v.CreatedDate)
	}

	mock.ExpectQuery(query).
		WithArgs(testUser.ID, testCreatedDate).
		WillReturnRows(rows)

	got, err := ldb.RetrieveTasks(ctx, ns(testUser.ID), ns(testCreatedDate))
	if assert.NoError(t, err) {
		assert.ElementsMatch(t, got, testTaskList)
	}
}

func TestRetrieveTaskEmpty(t *testing.T) {
	ldb, mock := newSQLMock()
	defer ldb.DB.Close()

	ctx := context.Background()
	query := `SELECT id, content, user_id, created_date FROM tasks WHERE user_id = (.+) AND created_date = (.+)`
	rows := sqlmock.NewRows([]string{"id", "content", "user_id", "created_date"})
	mock.ExpectQuery(query).WillReturnRows(rows)
	mock.ExpectQuery(query).WillReturnRows(rows)
	assert := assert.New(t)

	got, err := ldb.RetrieveTasks(ctx, ns("anotherUserID"), ns(testCreatedDate))
	if assert.NoError(err) {
		assert.Empty(got)
	}

	got, err = ldb.RetrieveTasks(ctx, ns(testUser.ID), ns("2020-11-22"))
	if assert.NoError(err) {
		assert.Empty(got)
	}
}

func TestAddTaskSuccess(t *testing.T) {
	ldb, mock := newSQLMock()
	defer ldb.DB.Close()

	ctx := context.Background()
	query := `^INSERT INTO tasks \(id, content, user_id, created_date\) `
	task := testTaskList[0]
	rows := sqlmock.NewResult(1, 1)
	mock.ExpectBegin()
	mock.ExpectExec(query).
		WithArgs(task.ID, task.Content, task.UserID, task.CreatedDate,
			task.UserID, task.CreatedDate, task.UserID).
		WillReturnResult(rows)
	mock.ExpectCommit()

	numRowAffected, err := ldb.AddTask(ctx, task)
	if assert.NoError(t, err) {
		assert.Equal(t, numRowAffected, int64(1))
	}
}
