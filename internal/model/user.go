package model

import "context"

type UserStore interface {
	Create(ctx context.Context, user *User) error
	FindUser(ctx context.Context, userID string) (*User, error)
}
