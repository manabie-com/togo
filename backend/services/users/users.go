package users

import (
	"github.com/jssoriao/todo-go/storage"
	"github.com/jssoriao/todo-go/storage/dynamodb"
)

var _ UsersStore = (*dynamodb.Storage)(nil)

type User struct {
	ID         string
	DailyLimit int
}

type UsersStore interface {
	CreateUser(storage.User) (storage.User, error)
	GetUser(string) (*storage.User, error)
}

type Handler struct {
	// TODO: Add logger
	store UsersStore
}

// NewHandler returns new handler for this service.
func NewHandler(store UsersStore) (*Handler, error) {
	return &Handler{
		store: store,
	}, nil
}
