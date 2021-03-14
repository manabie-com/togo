package domain

import "time"

// User reflects users data from DB
type User struct {
	ID        int
	Username  string
	Password  string
	CreatedAt *time.Time `db:"created_at"`
}

type UserAuthParam struct {
	Username string
	Password string
}

type UserRepository interface {
	GetByID(id int) (*User, error)
	GetByCredentials(username string, password string) (*User, error)
}
