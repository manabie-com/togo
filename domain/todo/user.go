package todo

import (
	"context"
	"fmt"

	"github.com/laghodessa/togo/domain"
)

// User of the todo context
// It contains user's settings
type User struct {
	ID string
	// TaskDailyLimit is the maximum tasks that can be added per day
	TaskDailyLimit int
}

// HitTaskDailyLimit returns error when the task daily limit has been reached
// and won't accept any more tasks
func (u User) HitTaskDailyLimit(todayTotal int) error {
	if todayTotal >= u.TaskDailyLimit {
		return ErrUserHitTaskDailyLimit
	}
	return nil
}

type UserRepo interface {
	// GetUser returns user by id
	GetUser(ctx context.Context, id string) (User, error)
	// AddUser creates a new user
	AddUser(context.Context, User) error
}

type UserOpt func(*User) error

func NewUser(opts ...UserOpt) (u User, err error) {
	for _, applyOpt := range opts {
		if err := applyOpt(&u); err != nil {
			return User{}, err
		}
	}

	if u.TaskDailyLimit == 0 {
		return User{}, fmt.Errorf("%w: missing task daily limit", domain.ErrInvalidArg)
	}
	u.ID = domain.NewID()
	return u, nil
}

func UserTaskDailyLimit(limit int) UserOpt {
	return func(u *User) error {
		if limit <= 0 {
			return domain.InvalidArg("invalid task daily limit")
		}
		u.TaskDailyLimit = limit
		return nil
	}
}

var (
	ErrUserHitTaskDailyLimit = domain.Error{
		Code:    "todo_user_hit_task_daily_limit",
		Message: "user hit task daily limit",
	}
)
