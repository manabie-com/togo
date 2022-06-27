package repository

import "time"

type UserRepository interface {
	IsUserHavingMaxTodo(userID int64, date time.Time) error
}
