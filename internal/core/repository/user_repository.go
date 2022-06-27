package repository

import "time"

type UserRepository interface {
	CountTodosByDay(userID int64, date time.Time) (int64, error)
}
