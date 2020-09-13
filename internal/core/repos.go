package core

import (
  "context"
  "errors"
  "time"
)

var ErrUserNotFound = errors.New("user not found")

type TaskRepo interface {
  Create(ctx context.Context, task *Task) error
  ByUser(ctx context.Context, userId string) ([]*Task, error)
  ByUserDate(ctx context.Context, userId string, date time.Time) ([]*Task, error)
}

type UserRepo interface {
  ById(ctx context.Context, id string) (*User, error)
  Validate(ctx context.Context, userId string, password string) (*User, bool)
}
