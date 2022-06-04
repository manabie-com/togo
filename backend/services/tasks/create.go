package tasks

import (
	"context"

	"github.com/jssoriao/todo-go/storage"
)

func (h *Handler) CreateTask(ctx context.Context, payload *Task) (*Task, error) {
	user, err := h.store.GetUser(payload.UserID)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, ErrUserNotFound
	}
	count, err := h.store.CountTasksForTheDay(payload.UserID, payload.DueDate)
	if err != nil {
		return nil, err
	}
	if count >= user.DailyLimit {
		return nil, ErrExceededTasksLimit
	}

	task, err := h.store.CreateTask(storage.Task{
		UserID:  payload.UserID,
		Title:   payload.Title,
		Done:    false,
		DueDate: payload.DueDate.Unix(),
	})
	if err != nil {
		return nil, err
	}
	return &Task{
		ID:      task.ID,
		UserID:  task.UserID,
		Title:   task.Title,
		Done:    task.Done,
		DueDate: payload.DueDate,
	}, nil
}
