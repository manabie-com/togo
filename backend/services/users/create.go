package users

import (
	"context"

	"github.com/jssoriao/todo-go/storage"
)

func (h *Handler) CreateUser(ctx context.Context, payload *User) (*User, error) {
	user, err := h.store.CreateUser(storage.User{DailyLimit: payload.DailyLimit})
	if err != nil {
		// TODO: Log the error
		return nil, err
	}
	return &User{
		ID:         user.ID,
		DailyLimit: user.DailyLimit,
	}, nil
}
