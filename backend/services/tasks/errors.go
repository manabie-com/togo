package tasks

import "errors"

var ErrUserNotFound = errors.New("user not found")
var ErrExceededTasksLimit = errors.New("user's count of tasks already exceeded limit")
