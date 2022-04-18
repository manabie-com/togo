package users

import (
	"testing"
	"time"

	"github.com/kozloz/togo"
	"github.com/kozloz/togo/internal/store/test"
)

var op *Operation

func TestMain(m *testing.M) {
	store := &test.Store{}
	op = NewOperation(store)
	m.Run()
}
func TestGet(t *testing.T) {
	userID := 1
	user, err := op.Get(int64(userID))
	if err != nil {
		t.Errorf("Failed to get user with error %v", err)
	}
	if user.ID != int64(userID) {
		t.Errorf("Failed user get, expected '%d', got '%d'", userID, user.ID)

	}

	user, err = op.Get(0)
	if err != nil {
		t.Errorf("Failed to get user with error %v", err)
	}
	if user != nil {
		t.Errorf("Failed user get, expected nil, got '%v'", user)

	}
}

func TestCreate(t *testing.T) {
	userID := 1
	user, err := op.Create(int64(userID))
	if err != nil {
		t.Errorf("Failed to create user with error %v", err)
	}
	if user.ID != int64(userID) {
		t.Errorf("Failed user create, expected '%d', got '%d'", userID, user.ID)
	}
}

func TestUpdate(t *testing.T) {
	userID := 1
	user, err := op.Get(int64(userID))
	if err != nil {
		t.Errorf("Failed to get user with error %v", err)
	}
	testTask := &togo.Task{
		ID:     2,
		UserID: int64(userID),
		Name:   "testtest",
	}
	tasks := append(user.Tasks, testTask)
	user.Tasks = tasks
	counter := &togo.DailyCounter{
		UserID:      user.ID,
		DailyCount:  5,
		LastUpdated: time.Now(),
	}
	user.DailyCounter = counter
	user, err = op.Update(user)
	if err != nil {
		t.Errorf("Failed to update user with error %v", err)
	}
	if len(user.Tasks) != len(tasks) {
		t.Errorf("Mismatch on updated tasks length, expected %d, got %d", len(tasks), len(user.Tasks))
	}
	if user.DailyCounter.DailyCount != counter.DailyCount {
		t.Errorf("Mismatch on updated daily count, expected %d, got %d", counter.DailyCount, user.DailyCounter.DailyCount)
	}

}
