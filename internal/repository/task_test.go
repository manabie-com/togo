package repository

import (
	e "lntvan166/togo/internal/entities"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

var task = &e.Task{
	ID:          1,
	Name:        "test",
	Description: "test",
	CreatedAt:   "2020-01-01",
	Completed:   false,
	UserID:      1,
}

func TestAddTask(t *testing.T) {
	db, mock := NewMock()
	repo := &repository{db}
	defer db.Close()

	query := regexp.QuoteMeta(`INSERT INTO tasks (
		name, description, created_at, completed, user_id)
		VALUES ($1, $2, $3, $4, $5);`)

	mock.ExpectBegin()
	mock.ExpectExec(query).WithArgs(task.Name, task.Description, task.CreatedAt, task.Completed, task.UserID).WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	err := repo.CreateTask(task)
	assert.NoError(t, err)
}

func TestGetAllTask(t *testing.T) {
	db, mock := NewMock()
	repo := &repository{db}
	defer db.Close()

	query := regexp.QuoteMeta(`SELECT * FROM tasks;`)

	rows := sqlmock.NewRows([]string{"id", "name", "description", "created_at", "completed", "user_id"}).
		AddRow(task.ID, task.Name, task.Description, task.CreatedAt, task.Completed, task.UserID)

	mock.ExpectQuery(query).WillReturnRows(rows)

	tasks, err := repo.GetAllTask()
	assert.NoError(t, err)
	assert.Equal(t, 1, len(*tasks))
}

func TestGetTaskByID(t *testing.T) {
	db, mock := NewMock()
	repo := &repository{db}
	defer db.Close()

	query := regexp.QuoteMeta(`SELECT * FROM tasks WHERE id = $1;`)

	rows := sqlmock.NewRows([]string{"id", "name", "description", "created_at", "completed", "user_id"}).
		AddRow(task.ID, task.Name, task.Description, task.CreatedAt, task.Completed, task.UserID)

	mock.ExpectQuery(query).WithArgs(task.ID).WillReturnRows(rows)

	newTask, err := repo.GetTaskByID(task.ID)
	assert.NoError(t, err)
	assert.Equal(t, task, newTask)
}

// TODO: Check limit task today
// func TestCheckLimitTaskToday(t *testing.T) {
// 	db, mock := NewMock()
// 	repo := &repository{db}
// 	defer db.Close()

// 	query := regexp.QuoteMeta(`SELECT COUNT(*) FROM tasks WHERE user_id = $1 AND created_at >= current_date;`)

// 	rows := sqlmock.NewRows([]string{"count"}).
// 		AddRow(1)

// 	mock.ExpectQuery(query).WithArgs(u.ID).WillReturnRows(rows)

// 	isLimit, err := repo.CheckLimitTaskToday(u.ID)
// 	assert.NoError(t, err)
// 	assert.Equal(t, false, isLimit)
// }

func TestUpdateTask(t *testing.T) {
	db, mock := NewMock()
	repo := &repository{db}
	defer db.Close()

	query := regexp.QuoteMeta(`UPDATE tasks SET name = $1, description = $2, created_at = $3, completed = $4, user_id = $5 WHERE id = $6;`)

	mock.ExpectBegin()
	mock.ExpectExec(query).WithArgs(task.Name, task.Description, task.CreatedAt, task.Completed, task.UserID, task.ID).WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	err := repo.UpdateTask(task)
	assert.NoError(t, err)
}

func TestCompleteTask(t *testing.T) {
	db, mock := NewMock()
	repo := &repository{db}
	defer db.Close()

	query := regexp.QuoteMeta(`UPDATE tasks SET completed = true WHERE id = $1;`)

	mock.ExpectBegin()
	mock.ExpectExec(query).WithArgs(task.ID).WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	err := repo.CompleteTask(task.ID)
	assert.NoError(t, err)
}

func TestDeleteTask(t *testing.T) {
	db, mock := NewMock()
	repo := &repository{db}
	defer db.Close()

	query := regexp.QuoteMeta(`DELETE FROM tasks WHERE id = $1;`)

	mock.ExpectBegin()
	mock.ExpectExec(query).WithArgs(task.ID).WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	err := repo.DeleteTask(task.ID)
	assert.NoError(t, err)
}

func TestDeleteAllTask(t *testing.T) {
	db, mock := NewMock()
	repo := &repository{db}
	defer db.Close()

	query := regexp.QuoteMeta(`DELETE FROM tasks;`)

	mock.ExpectBegin()
	mock.ExpectExec(query).WithArgs().WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	err := repo.DeleteAllTask()
	assert.NoError(t, err)
}

func TestDeleteAllTaskOfUser(t *testing.T) {
	db, mock := NewMock()
	repo := &repository{db}
	defer db.Close()

	query := regexp.QuoteMeta(`DELETE FROM tasks WHERE user_id = $1;`)

	mock.ExpectBegin()
	mock.ExpectExec(query).WithArgs(u.ID).WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	err := repo.DeleteAllTaskOfUser(u.ID)
	assert.NoError(t, err)
}
