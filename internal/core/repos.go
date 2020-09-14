package core

import (
  "context"
  "errors"
  "time"
)

var (
  ErrUserNotFound = errors.New("user not found")
  ErrUserAlreadyExists = errors.New("user already exists")
  ErrWrongIdPassword = errors.New("invalid user id or password")
  ErrEmptyId = errors.New("empty user id")
  ErrEmptyPassword = errors.New("empty password")
  ErrInvalidMaxTodo = errors.New("invalid max todo")
)

type TaskRepo interface {
  Create(ctx context.Context, task *Task) error
  ByUser(ctx context.Context, userId string) ([]*Task, error)
  ByUserDate(ctx context.Context, userId string, date time.Time) ([]*Task, error)
}

type Hasher interface {
  Hash(password string) (string, error)
  Compare(password, hash string) bool
}

type UserRepo interface {
  Hasher

  Create(ctx context.Context, user *User, password string) error
  ById(ctx context.Context, id string) (*User, error)
  Validate(ctx context.Context, userId string, password string) (*User, error)
}
