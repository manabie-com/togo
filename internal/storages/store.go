package storages

import "context"

type Store interface {
	InitTables() error
	RetrieveTasks(ctx context.Context, userID, createdDate string) ([]*Task, error)
	AddTask(ctx context.Context, t *Task) (canAdd bool, err error)
	AddUser(ctx context.Context, user *User) error
	SetUserPassword(ctx context.Context, id, password string) error
	MaxTodo(ctx context.Context, userID string) (int, error)
	CanUserCreateTodo(ctx context.Context, t *Task) (bool, error)
	ValidateUser(ctx context.Context, userID, pwd string) bool
}
