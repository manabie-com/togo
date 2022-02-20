package domain

import (
	"context"
	"errors"
)

var (
	// ErrUserNotFound user not found error
	ErrUserNotFound = errors.New("USER_NOT_FOUND")
	// ErrUserCreateFailed failed to create a user
	ErrUserCreateFailed = errors.New("USER_CREATE_FAILED")
	// ErrDuplicatedUsername duplicate username
	ErrDuplicatedUsername = errors.New("DUPLICATED_USERNAME")
)

// User model
type User struct {
	ID       uint   `json:"id,omitempty" gorm:"primarykey"`
	FullName string `json:"fullName,omitempty"`
	Username string `json:"username,omitempty" gorm:"uniqueIndex"`
	Password string `json:"-"`

	TasksPerDay int `json:"tasksPerDay"`
}

// UserService service interface
type UserService interface {
	CreateUser(ctx context.Context, user *User) (*User, error)
	GetUserByID(ctx context.Context, id uint) (*User, error)
	UpdateByID(ctx context.Context, id uint, update *User) (*User, error)
}
