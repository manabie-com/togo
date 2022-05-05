package models

import (
	"errors"
	"github.com/jmsemira/togo/internal/database"
	"time"
)

type Todo struct {
	ID        uint
	Name      string
	UserID    uint
	CreatedAt *time.Time
}

func (t *Todo) Save(userLimit int) error {
	// check if User exceed the ratelimit
	db := database.GetDB()

	now := time.Now()
	todos := []Todo{}

	db.Where("created_at > (?) and user_id = ?", now.Format("2006-01-02"), t.UserID).Find(&todos)

	// if user limit is set to 0 there will be no limit
	if len(todos) < userLimit || userLimit == 0 {
		db.Create(t)
		if t.ID == 0 {
			return errors.New("Error saving todo")
		}
		return nil
	}
	return errors.New("Rate Limit for today was reached")
}
