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
	// ErrInvalidUser invalid user error
	ErrInvalidUser = errors.New("INVALID_USER")
)

// User model
type User struct {
	ID       uint   `json:"id" gorm:"primarykey"`
	FullName string `json:"fullName"`
	Username string `json:"username" gorm:"uniqueIndex"`
	Password string `json:"-"`

	TasksPerDay int `json:"tasksPerDay"`
}

// Validate valiate user data
func (m User) Validate() error {
	if m.Username == "" || len(m.Password) < 6 {
		return ErrInvalidUser
	}
	return nil
}

// UserService service interface
type UserService interface {
	CreateUser(ctx context.Context, user *User) (*User, error)
	GetUserByID(ctx context.Context, id uint) (*User, error)
	UpdateByID(ctx context.Context, id uint, update *User) (*User, error)
}
