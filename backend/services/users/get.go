package users

import (
	"context"
)

func (h *Handler) GetUser(ctx context.Context, id string) (*User, error) {
	user, err := h.store.GetUser(id)
	if err != nil {
		// TODO: Log the error
		return nil, err
	}
	if user == nil {
		return nil, ErrUserNotFound
	}
	return &User{
		ID:         user.ID,
		DailyLimit: user.DailyLimit,
	}, nil
}
