package tasks

import (
	"testing"

	"github.com/kozloz/togo/internal/errors"
	"github.com/kozloz/togo/internal/store/test"
	"github.com/kozloz/togo/internal/users"
)

var op *Operation

func TestMain(m *testing.M) {
	store := &test.Store{}
	userOp := users.NewOperation(store)
	op = NewOperation(store, userOp)
	m.Run()
}
func TestCreate(t *testing.T) {
	userID := 1
	task1 := "task1"
	task, err := op.Create(int64(userID), task1)
	if err != nil {
		t.Errorf("Failed to create task, got error %v", err)
	}
	if task.UserID != int64(userID) {
		t.Errorf("Failed create task, expected user ID '%d', got '%d'", userID, task.UserID)
	}
	if task.Name != task1 {
		t.Errorf("Failed create task, expected task name '%s', got '%s'", task1, task.Name)
	}
	user, _ := op.userOperation.Get(int64(userID))
	totalTasks := 1
	if len(user.Tasks) != totalTasks {
		t.Errorf("Expected len user tasks to be %d, got %d", totalTasks, len(user.Tasks))
	}

	expectedDailyCounter := 1
	if user.DailyCounter.DailyCount != expectedDailyCounter {
		t.Errorf("Expected daily counter to be %d, got %d", expectedDailyCounter, user.DailyCounter.DailyCount)
	}

	// Test that a user can create more than one task
	task2 := "task2"
	_, err = op.Create(int64(userID), task2)
	if err != nil {
		t.Errorf("Failed to create task 2, got error %v", err)
	}
	user, _ = op.userOperation.Get(int64(userID))
	totalTasks = 2
	if len(user.Tasks) != totalTasks {
		t.Errorf("Expected len user tasks to be %d, got %d", totalTasks, len(user.Tasks))
	}
	expectedDailyCounter = 2
	if user.DailyCounter.DailyCount != expectedDailyCounter {
		t.Errorf("Expected daily counter to be %d, got %d", expectedDailyCounter, user.DailyCounter.DailyCount)
	}
	// Test that a user can't create more than the user's daily limit(3)
	task3 := "task3"
	task, err = op.Create(int64(userID), task3)
	if err != errors.MaxLimit {
		t.Errorf("Expected %v, got %v", errors.MaxLimit, err)
	}
	if task != nil {
		t.Errorf("Expected nil, got %v", task)
	}
	user, _ = op.userOperation.Get(int64(userID))
	totalTasks = 2
	if len(user.Tasks) != totalTasks {
		t.Errorf("Expected len user tasks to be %d, got %d", totalTasks, len(user.Tasks))
	}
	expectedDailyCounter = 2
	if user.DailyCounter.DailyCount != expectedDailyCounter {
		t.Errorf("Expected daily counter to be %d, got %d", expectedDailyCounter, user.DailyCounter.DailyCount)
	}

	// Test that a user that hasn't created a task today is able to
	userID2 := 2
	task4 := "task4"
	task, err = op.Create(int64(userID2), task4)
	if err != nil {
		t.Errorf("Failed to create task, got error %v", err)
	}
	if task.UserID != int64(userID2) {
		t.Errorf("Failed create task, expected user ID '%d', got '%d'", userID2, task.UserID)
	}
	if task.Name != task4 {
		t.Errorf("Failed create task, expected task name '%s', got '%s'", task4, task.Name)
	}

	user2, _ := op.userOperation.Get(int64(userID2))
	totalTasks = 1
	if len(user2.Tasks) != totalTasks {
		t.Errorf("Expected len user tasks to be %d, got %d", totalTasks, len(user2.Tasks))
	}
	expectedDailyCounter = 1
	if user2.DailyCounter.DailyCount != expectedDailyCounter {
		t.Errorf("Expected daily counter to be %d, got %d", expectedDailyCounter, user2.DailyCounter.DailyCount)
	}
}
