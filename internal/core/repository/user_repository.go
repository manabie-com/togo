package repository

import (
	"time"
)

type UserRepository interface {
	IsUserExisted(userID int64) error
	IsUserHavingMaxTodo(userID int64, date time.Time) error
}
