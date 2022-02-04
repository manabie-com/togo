package togo

import "fmt"

// User represents a user
type User struct {
	ID         int64
	Username   string // Username is just for display purposes
	Tasks      []Task
	DailyLimit int
	DailyCount int
}

func (u *User) String() string {
	return fmt.Sprintf("ID: '%d', Username: '%s', DailyLimit: '%d', Tasks: '%v'",
		u.ID, u.Username, u.DailyLimit, u.Tasks)
}
