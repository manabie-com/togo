package postgres

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

	retrieveTasks := "SELECT id, content, user_id FROM task WHERE user_id = \\$1 AND deleted_at IS NULL"
	rows := sqlmock.NewRows([]string{"id", "content", "user_id"}).
		AddRow(mockTask.ID, mockTask.Content, mockTask.UserID)

	mock.ExpectQuery(retrieveTasks).WithArgs(mockTask.ID).WillReturnRows(rows)

	task, err := repo.RetrieveTasks(context.Background(), mockTask.ID)
	if err != nil {
		t.Errorf("Data Error")
	}
	if !reflect.DeepEqual(task, []taskEntity.Task{mockTask}) {
		t.Errorf("response not match")
	}
}

func TestRetrieveTasksErr(t *testing.T) {
	db, mock := NewMock()
	repo := &taskRepository{db}
	defer func() {
		repo.Close()
	}()

	retrieveTasks := "SELECT id, content, user_id FROM task WHERE user_id = \\$1 AND deleted_at IS NULL"
	mock.ExpectQuery(retrieveTasks).WithArgs(mockTask.ID).WillReturnError(errors.New("DatabaseError"))

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

	addTask := "INSERT INTO task \\(id, content, user_id, created_at\\) VALUES \\(\\$1, \\$2, \\$3, \\$4\\)"
	now := time.Now().UTC()
	mock.ExpectExec(addTask).WithArgs(mockTask.ID, mockTask.Content, mockTask.UserID, now).WillReturnResult(sqlmock.NewResult(0, 1))


	err := repo.AddTask(context.Background(), mockTask, now)
	if err != nil {
		t.Errorf("Data Error")
	}
}

func TestCreateTaskErr(t *testing.T) {
	db, mock := NewMock()
	repo := &taskRepository{db}
	defer func() {
		repo.Close()
	}()

	addTask := "INSERT INTO task \\(id, content, user_id, created_at\\) VALUES \\(\\$1, \\$2, \\$3, \\$4\\)"
	now := time.Now().UTC()
	mock.ExpectExec(addTask).WithArgs(mockTask.ID, mockTask.Content, mockTask.UserID, now).WillReturnError(errors.New("error"))

	err := repo.AddTask(context.Background(), mockTask, now)
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

	now := time.Now().UTC()
	softDeleteTask := "UPDATE task SET deleted_at = \\$1 WHERE id = \\$2"
	mock.ExpectExec(softDeleteTask).WithArgs(now, mockTask.ID).WillReturnResult(sqlmock.NewResult(0, 1))

	err := repo.SoftDeleteTask(context.Background(), mockTask.ID, now)
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

	softDeleteTask := "UPDATE task SET deleted_at = \\$1 WHERE id = \\$2"
	now := time.Now().UTC()
	mock.ExpectExec(softDeleteTask).WithArgs(now, mockTask.ID).WillReturnError(errors.New("error"))

	err := repo.SoftDeleteTask(context.Background(), mockTask.ID, now)
	if err == nil {
		t.Errorf("Soft Delete Task Failed")
	}
}