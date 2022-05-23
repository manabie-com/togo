package repository_test

import (
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"testing"
	"time"
	"togo/domain"
	taskRepository "togo/internal/task/repository"
)

func NewMock() (*gorm.DB, sqlmock.Sqlmock) {
	testDB, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		log.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	db, dbErr := gorm.Open(mysql.New(mysql.Config{
		Conn:                      testDB,
		SkipInitializeWithVersion: true,
	}), &gorm.Config{})
	if dbErr != nil {
		log.Fatalf("has error '%s' to connect test database", dbErr.Error())
	}

	return db, mock
}

func TestCreate(t *testing.T) {
	now := time.Now()
	userId := 1
	task := domain.Task{
		Content:   "This is a task",
		UserId:    &userId,
		CreatedAt: now,
	}
	db, mock := NewMock()

	mock.ExpectBegin()
	mock.ExpectExec("INSERT INTO `tasks` (`content`,`user_id`,`created_at`) VALUES (?,?,?)").
		WithArgs(task.Content, *task.UserId, task.CreatedAt).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	taskRepo := taskRepository.NewTaskRepository(db)

	_, createErr := taskRepo.Create(task)
	assert.NoError(t, createErr)
}

func TestUpdate(t *testing.T) {
	now := time.Now()
	userId := 22
	task := domain.Task{
		ID:        1,
		Content:   "task content",
		UserId:    &userId,
		CreatedAt: now,
	}

	db, mock := NewMock()

	mock.ExpectBegin()
	mock.ExpectExec("UPDATE `tasks` SET `content`=?,`user_id`=?,`created_at`=? WHERE `id` = ?").
		WithArgs(task.Content, *task.UserId, task.CreatedAt, task.ID).
		WillReturnResult(sqlmock.NewResult(0, 1))
	mock.ExpectCommit()

	taskRepo := taskRepository.NewTaskRepository(db)
	updateErr := taskRepo.Update(task)
	assert.NoError(t, updateErr)
}
