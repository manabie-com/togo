package tasks

import (
	"time"

	"github.com/jssoriao/todo-go/storage"
)

type Task struct {
	ID      string
	UserID  string
	Title   string
	Done    bool
	DueDate time.Time
}

type TasksStore interface {
	CreateTask(storage.Task) (storage.Task, error)
	CountTasksForTheDay(userID string, dueDate time.Time) (int, error)
	GetUser(id string) (*storage.User, error)
}

type Handler struct {
	// TODO: Add logger
	store TasksStore
}

// NewHandler returns new handler for this service.
func NewHandler(store TasksStore) (*Handler, error) {
	return &Handler{
		store: store,
	}, nil
}
