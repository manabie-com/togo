package users_test

import (
	"context"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	"github.com/jssoriao/todo-go/services/users"
	"github.com/jssoriao/todo-go/storage"
)

type mockUsersStore struct{}

func (m mockUsersStore) CreateUser(user storage.User) (storage.User, error) {
	timestamp := time.Now()
	user.ID = "userId"
	user.Created = timestamp
	user.Updated = timestamp
	return user, nil
}

func (m mockUsersStore) GetUser(string) (*storage.User, error) {
	// Unimplemented
	return nil, nil
}

func TestCreateUser(t *testing.T) {
	store := mockUsersStore{}
	h, _ := users.NewHandler(store)
	task, err := h.CreateUser(context.Background(), &users.User{
		DailyLimit: 10,
	})
	if err != nil {
		t.Errorf("error must be nil")
	}
	if diff := cmp.Diff(task, &users.User{
		ID:         "userId",
		DailyLimit: 10,
	}); diff != "" {
		t.Errorf("result mismatch %s", diff)
	}
}
