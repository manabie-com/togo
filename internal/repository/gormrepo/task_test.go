package gormrepo

import (
	"context"
	"errors"
	"regexp"
	"testing"
	"togo/internal/domain"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func Test_taskRepository_Create_failed(t *testing.T) {
	taskInput := &domain.Task{
		UserID:  1,
		Content: "text",
	}
	sqlErr := errors.New("insert failed")
	db, mock, gdb, err := newGormDBMock()
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()
	mock.ExpectBegin()
	mock.ExpectQuery("INSERT INTO \"tasks\"").WithArgs(1, "text").WillReturnError(sqlErr)
	mock.ExpectRollback()
	r := taskRepository{gdb}
	task, err := r.Create(context.Background(), taskInput)
	assert.Nil(t, task)
	assert.ErrorIs(t, err, sqlErr)
}

func Test_taskRepository_Create_Successful(t *testing.T) {
	taskInput := &domain.Task{
		UserID:  1,
		Content: "text",
	}
	eTask := &domain.Task{
		ID:      1,
		UserID:  1,
		Content: "text",
	}
	db, mock, gdb, err := newGormDBMock()
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()
	columns := []string{"id", "user_id", "content"}
	row := sqlmock.NewRows(columns).AddRow(
		eTask.ID,
		eTask.UserID,
		eTask.Content,
	)
	mock.ExpectBegin()
	mock.ExpectQuery("INSERT INTO \"tasks\"").WillReturnRows(row)
	mock.ExpectCommit()
	r := taskRepository{gdb}
	task, err := r.Create(context.Background(), taskInput)
	assert.NoError(t, err)
	assert.Equal(t, eTask, task)
}

func Test_taskRepository_Update_NotFound(t *testing.T) {
	filter := &domain.Task{
		ID:     1,
		UserID: 1,
	}
	update := &domain.Task{
		Content: "Updated",
	}
	sqlErr := errors.New("insert failed")
	db, mock, gdb, err := newGormDBMock()
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()
	mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM \"tasks\"")).WillReturnError(sqlErr)
	r := taskRepository{gdb}
	_, err = r.Update(context.Background(), filter, update)
	assert.ErrorIs(t, err, sqlErr)
}

func Test_taskRepository_Update_Successful(t *testing.T) {
	filter := &domain.Task{
		ID:     1,
		UserID: 1,
	}
	eTask := &domain.Task{
		ID:      1,
		UserID:  1,
		Content: "updated",
	}
	update := &domain.Task{
		Content: "updated",
	}
	db, mock, gdb, err := newGormDBMock()
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()
	row := sqlmock.NewRows([]string{"id", "user_id", "content"}).AddRow(1, 1, "text")
	mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM \"tasks\"")).WillReturnRows(row)
	mock.ExpectBegin()
	mock.ExpectExec(regexp.QuoteMeta("UPDATE \"tasks\"")).WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()
	r := taskRepository{gdb}
	task, err := r.Update(context.Background(), filter, update)
	assert.NoError(t, err)
	assert.Equal(t, eTask, task)
}

func Test_taskRepository_Find_NotFound(t *testing.T) {
	filter := &domain.Task{
		UserID: 1,
	}
	eTasks := make([]*domain.Task, 0)
	rows := sqlmock.NewRows([]string{"id", "user_id", "content"})
	db, mock, gdb, err := newGormDBMock()
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()
	mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM \"tasks\"")).WillReturnRows(rows)
	r := taskRepository{gdb}
	tasks, err := r.Find(context.Background(), filter)
	assert.NoError(t, err)
	assert.Equal(t, eTasks, tasks)
}

func Test_taskRepository_Find_Successful(t *testing.T) {
	filter := &domain.Task{
		UserID: 1,
	}
	eTasks := []*domain.Task{
		{
			ID:      1,
			UserID:  1,
			Content: "text 1",
		},
		{
			ID:      2,
			UserID:  1,
			Content: "text 2",
		},
		{
			ID:      3,
			UserID:  1,
			Content: "text 3",
		},
	}
	rows := sqlmock.NewRows([]string{"id", "user_id", "content"})
	for _, task := range eTasks {
		rows.AddRow(task.ID, task.UserID, task.Content)
	}
	db, mock, gdb, err := newGormDBMock()
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()
	mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM \"tasks\"")).WillReturnRows(rows)
	r := taskRepository{gdb}
	tasks, err := r.Find(context.Background(), filter)
	assert.NoError(t, err)
	assert.Equal(t, eTasks, tasks)
}
