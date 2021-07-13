package task

import (
	"errors"
	"fmt"
)

var ErrReachedOutTaskTodoPerDay = errors.New("task todo is enough for today")

type UserNotExistedError struct {
	userID string
}

func NewUserNotExistedError(userID string) error {
	return &UserNotExistedError{
		userID: userID,
	}
}

func (e UserNotExistedError) Error() string {
	return fmt.Sprintf("user is not existed with user_id=%s", e.userID)
}
