package test

import (
	"context"
	"github.com/google/uuid"
	"github.com/manabie-com/togo/internal/models"
	"github.com/manabie-com/togo/internal/repositories"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"reflect"
	"testing"
)

var userId = "123"
var createdDate = "2020-10-10"
var taskRepository repositories.TaskRepository

func init() {
	db, err := gorm.Open(sqlite.Open("../../../data.db"), &gorm.Config{})
	if err != nil {
		logrus.Errorf("Connect to DB error: %s", err.Error())
		panic(err)
	}

	taskRepository = repositories.NewTaskRepository(db)
}

func TestCreate(t *testing.T) {
	newTask := &models.Task{
		ID:          uuid.New().String(),
		Content:     "Test-New-Task",
		UserID:      userId,
		CreatedDate: createdDate,
	}

	var expectedTask = newTask
	var expectedErr error = nil

	got, err := taskRepository.Create(context.Background(), newTask)
	if err != expectedErr {
		t.Errorf("Expected nil error, but got err: %v", err)
	}

	if !reflect.DeepEqual(expectedTask, got) {
		t.Errorf("Expected task %v, but got %v", expectedTask, got)
	}
}

func TestGetTasks(t *testing.T) {
	tasks, err := taskRepository.GetTasks(context.Background(), userId, createdDate)
	var expectedErr error = nil
	if err != expectedErr {
		t.Errorf("Expected nil error, but got err: %v", err)
	}

	if len(tasks) < 1 {
		t.Error("Expected having tasks but got 0")
	}
}
