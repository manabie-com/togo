package entity

import (
	"errors"
	"time"
)

type Todo struct {
	UserID    int64
	Name      string
	Content   string
	CreatedAt time.Time
}

func (t Todo) Validate() error {
	if t.UserID <= 0 {
		return errors.New("invalid user_id")
	}
	if len(t.Name) == 0 {
		return errors.New("name is empty")
	}
	return nil
}
