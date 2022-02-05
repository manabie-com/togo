package togo

import (
	"fmt"
	"time"
)

// User represents a user
type User struct {
	ID         int64
	Tasks      []*Task
	DailyLimit int
	*DailyCounter
}

type DailyCounter struct {
	UserID      int64
	DailyCount  int
	LastUpdated time.Time
}

func (u *User) String() string {
	dailyCounter := 0
	if u.DailyCounter != nil {
		dailyCounter = u.DailyCounter.DailyCount
	}
	return fmt.Sprintf("ID: '%d', DailyCounter: '%d', DailyLimit: '%d', Tasks: '%v'",
		u.ID, dailyCounter, u.DailyLimit, u.Tasks)
}
