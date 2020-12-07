package sqllite

import (
	"context"
	"database/sql"
	"errors"
	"github.com/DATA-DOG/go-sqlmock"
	taskEntity "github.com/HoangVyDuong/togo/internal/storages/task"
	"reflect"
	"testing"
	"time"
)

var mockTask = taskEntity.Task{
	ID: 1234,
	Content: "content",
	UserID: 111,
}

func NewMock() (*sql.DB, sqlmock.Sqlmock) {
	db, mock, err := sqlmock.New()
	if err != nil {
		panic("error happen when new mock")
	}

	return db, mock
}

func TestRetrieveTasks(t *testing.T) {
	db, mock := NewMock()
	repo := &taskRepository{db}
	defer func() {
		repo.Close()
	}()

	query := "SELECT id, content, user_id FROM task WHERE user_id = \\? AND deleted_at IS NULL"

	rows := sqlmock.NewRows([]string{"id", "content", "user_id"}).
		AddRow(mockTask.ID, mockTask.Content, mockTask.UserID)

	mock.ExpectQuery(query).WithArgs(mockTask.ID).WillReturnRows(rows)

	task, err := repo.RetrieveTasks(context.Background(), mockTask.ID)
	if err != nil {
		t.Errorf("Data Error")
	}
	if !reflect.DeepEqual(task, mockTask) {
		t.Errorf("")
	}
}

func TestRetrieveTasksErr(t *testing.T) {
	db, mock := NewMock()
	repo := &taskRepository{db}
	defer func() {
		repo.Close()
	}()

	query := "SELECT id, content, user_id FROM task WHERE user_id = \\? WHERE deleted_at IS NULL"

	mock.ExpectQuery(query).WithArgs(mockTask.ID).WillReturnError(errors.New("DatabaseError"))

	_, err := repo.RetrieveTasks(context.Background(), mockTask.ID)
	if err == nil {
		t.Errorf("Data Error")
	}
}

func TestCreateTask(t *testing.T) {
	db, mock := NewMock()
	repo := &taskRepository{db}
	defer func() {
		repo.Close()
	}()

	query := "INSERT INTO task \\(id, content, user_id, created_at\\) VALUES (?, ?, ?, ?)"
	mock.ExpectExec(query).WithArgs(mockTask.ID, mockTask.Content, mockTask.UserID, time.Now()).WillReturnResult(sqlmock.NewResult(mockTask.ID ,1))

	taskID, err := repo.AddTask(context.Background(), mockTask)
	if err != nil {
		t.Errorf("Data Error")
	}
	if taskID != mockTask.ID {
		t.Errorf("Create Task Failed")
	}
}

func TestCreateTaskErr(t *testing.T) {
	db, mock := NewMock()
	repo := &taskRepository{db}
	defer func() {
		repo.Close()
	}()

	query := "INSERT INTO task \\(id, content, user_id, created_at\\) VALUES (?, ?, ?, ?)"
	mock.ExpectExec(query).WithArgs(mockTask.ID, mockTask.Content, mockTask.UserID, time.Now()).WillReturnResult(sqlmock.NewResult(mockTask.ID ,0))

	_, err := repo.AddTask(context.Background(), mockTask)
	if err == nil {
		t.Errorf("Data Error")
	}
}

func TestSoftDeleteTask(t *testing.T) {
	db, mock := NewMock()
	repo := &taskRepository{db}
	defer func() {
		repo.Close()
	}()

	query := "UPDATE task SET deleted_at = ? WHERE id = ?"
	mock.ExpectExec(query).WithArgs(mockTask.ID).WillReturnResult(sqlmock.NewResult(mockTask.ID, 1))

	err := repo.SoftDeleteTask(context.Background(), mockTask.ID)
	if err != nil {
		t.Errorf("Soft Delete Task Failed")
	}
}

func TestSoftDeleteTaskErr(t *testing.T) {
	db, mock := NewMock()
	repo := &taskRepository{db}
	defer func() {
		repo.Close()
	}()

	query := "UPDATE task SET deleted_at = ? WHERE id = ?"
	mock.ExpectExec(query).WithArgs(mockTask.ID).WillReturnResult(sqlmock.NewResult(mockTask.ID ,0))

	err := repo.SoftDeleteTask(context.Background(), mockTask.ID)
	if err == nil {
		t.Errorf("Soft Delete Task Failed")
	}
}