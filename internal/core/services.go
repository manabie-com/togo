package core

import (
  "context"
  "errors"
  "time"
)

var ErrMaxTodoReached = errors.New("user's maximum number of todo fixedDay reached")

type TaskService interface {
  IndexAll(ctx context.Context, user *User) ([]*Task, error)
  IndexDate(ctx context.Context, user *User, date time.Time) ([]*Task, error)
  Create(ctx context.Context, user *User, task *Task) error
  Update(ctx context.Context, user *User, task *Task) error
  Delete(ctx context.Context, user *User, id string) error
}

type UserService interface {
  Login(ctx context.Context, id, password string) (string, error)
  Signup(ctx context.Context, user *User, password string) (string, error)
}
