package sqllite

import (
	"context"
	"database/sql"
	"errors"
	"log"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/google/uuid"
	"github.com/manabie-com/togo/internal/storages"
	"github.com/manabie-com/togo/internal/util"
	"github.com/stretchr/testify/assert"
)

var task = storages.Task{
	ID:          util.ConvertSQLNullString(uuid.New().String()),
	Content:     util.ConvertSQLNullString("example"),
	UserID:      util.ConvertSQLNullString("firstUser"),
	CreatedDate: util.ConvertSQLNullString("2021-08-23"),
}

var user = storages.User{
	ID:       util.ConvertSQLNullString("firstUser"),
	Password: util.ConvertSQLNullString("example"),
	MaxTodo:  5,
}

func NewMock() (*sql.DB, sqlmock.Sqlmock) {
	db, mock, err := sqlmock.New()
	if err != nil {
		log.Printf("sqlmock db error %v", err)
	}
	return db, mock
}

func TestNewLiteDB(t *testing.T) {
	db, _ := NewMock()
	liteDB := &liteDB{db}

	newDB := NewLiteDB(db)
	assert.Equal(t, liteDB, newDB)
}

func TestRetrieveTasksSuccess(t *testing.T) {
	db, mock := NewMock()
	liteDB := &liteDB{db}

	stmt := "SELECT id, content, user_id, created_date FROM tasks WHERE user_id = \\? AND created_date = \\?"
	wantRows := sqlmock.NewRows([]string{"id", "content", "user_id", "created_date"}).
		AddRow(task.ID, task.Content, task.UserID, task.CreatedDate)

	mock.ExpectQuery(stmt).WithArgs(task.UserID, task.CreatedDate).WillReturnRows(wantRows)

	rows, err := liteDB.RetrieveTasks(context.Background(), task.UserID.String, task.CreatedDate.String)
	assert.NotNil(t, rows)
	assert.NoError(t, err)
}

func TestRetrieveTasksErrorQuery(t *testing.T) {
	db, mock := NewMock()
	liteDB := &liteDB{db}

	stmt := "SELECT id, content, user_id, created_date FROM tasksabc WHERE user_id = \\? AND created_date = \\?"
	wantRows := sqlmock.NewRows([]string{"id", "content", "user_id", "created_date"})

	mock.ExpectQuery(stmt).WithArgs(task.UserID, task.CreatedDate).WillReturnRows(wantRows)

	rows, err := liteDB.RetrieveTasks(context.Background(), task.UserID.String, task.CreatedDate.String)
	assert.Empty(t, rows)
	assert.Error(t, err)
}

func TestRetrieveTasksErrorScan(t *testing.T) {
	db, mock := NewMock()
	liteDB := &liteDB{db}

	stmt := "SELECT id, content, user_id, created_date FROM tasks WHERE user_id = \\? AND created_date = \\?"
	wantRows := sqlmock.NewRows([]string{"id", "content", "user_id"}).
		AddRow(task.ID, task.Content, task.UserID)

	mock.ExpectQuery(stmt).WithArgs(task.UserID, task.CreatedDate).WillReturnRows(wantRows)

	rows, err := liteDB.RetrieveTasks(context.Background(), task.UserID.String, task.CreatedDate.String)
	assert.Empty(t, rows)
	assert.Error(t, err)
}

func TestRetrieveTasksRowsError(t *testing.T) {
	db, mock := NewMock()
	liteDB := &liteDB{db}

	stmt := "SELECT id, content, user_id, created_date FROM tasks WHERE user_id = \\? AND created_date = \\?"
	wantRows := sqlmock.NewRows([]string{"id", "content", "user_id", "created_date"}).
		AddRow(task.ID, task.Content, task.UserID, task.CreatedDate).
		RowError(0, errors.New("rows err"))

	mock.ExpectQuery(stmt).WithArgs(task.UserID, task.CreatedDate).WillReturnRows(wantRows)

	rows, err := liteDB.RetrieveTasks(context.Background(), task.UserID.String, task.CreatedDate.String)
	assert.Empty(t, rows)
	assert.Error(t, err)
}

func TestAddTaskSuccess(t *testing.T) {
	db, mock := NewMock()
	liteDB := &liteDB{db}

	stmt := "INSERT INTO tasks \\(id, content, user_id, created_date\\) VALUES \\(\\?, \\?, \\?, \\?\\)"

	mock.ExpectExec(stmt).WithArgs(task.ID.String, task.Content.String, task.UserID.String, task.CreatedDate.String).WillReturnResult(sqlmock.NewResult(0, 1))

	err := liteDB.AddTask(context.Background(), task)
	assert.NoError(t, err)
}

func TestAddTaskError(t *testing.T) {
	db, mock := NewMock()
	liteDB := &liteDB{db}

	stmt := "INSERT INTO tasks \\(id, content, user_id, created_date\\) VALUES \\(\\?, \\?, \\?, \\?\\)"

	mock.ExpectExec(stmt).WithArgs(task.ID.String, task.Content.String, task.UserID.String, nil).WillReturnResult(sqlmock.NewResult(0, 0))

	err := liteDB.AddTask(context.Background(), task)
	assert.Error(t, err)
}

func TestValidateUserTrue(t *testing.T) {
	db, mock := NewMock()
	liteDB := &liteDB{db}

	stmt := "SELECT id FROM users WHERE id = \\? AND password = \\?"
	wantRows := sqlmock.NewRows([]string{"id"}).
		AddRow(user.ID)

	mock.ExpectQuery(stmt).WithArgs(user.ID.String, user.Password.String).WillReturnRows(wantRows)

	result := liteDB.ValidateUser(context.Background(), user.ID.String, user.Password.String)
	assert.True(t, result)
}

func TestValidateUserFalse(t *testing.T) {
	db, mock := NewMock()
	liteDB := &liteDB{db}

	stmt := "SELECT id FROM users WHERE id = \\? AND password = \\?"
	wantRows := sqlmock.NewRows([]string{"id"})

	mock.ExpectQuery(stmt).WithArgs(user.ID.String, user.Password.String).WillReturnRows(wantRows)

	result := liteDB.ValidateUser(context.Background(), user.ID.String, user.Password.String)
	assert.False(t, result)
}

func TestGetMaxTodoSuccess(t *testing.T) {
	db, mock := NewMock()
	liteDB := &liteDB{db}

	stmt := "SELECT max_todo FROM users WHERE id = \\?"
	wantRows := sqlmock.NewRows([]string{"id"}).
		AddRow(user.MaxTodo)

	mock.ExpectQuery(stmt).WithArgs(user.ID.String).WillReturnRows(wantRows)

	maxTodo, err := liteDB.GetMaxTodo(context.Background(), user.ID.String)
	assert.Equal(t, maxTodo, user.MaxTodo)
	assert.NoError(t, err)
}

func TestGetMaxTodoError(t *testing.T) {
	db, mock := NewMock()
	liteDB := &liteDB{db}

	stmt := "SELECT max_todo FROM users WHERE id = \\?"
	wantRows := sqlmock.NewRows([]string{"id"})

	mock.ExpectQuery(stmt).WithArgs(user.ID.String).WillReturnRows(wantRows)

	_, err := liteDB.GetMaxTodo(context.Background(), user.ID.String)
	assert.Error(t, err)
}
