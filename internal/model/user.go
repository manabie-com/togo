package model

import (
	"context"
)

// User reflects users data from DB
type User struct {
	ID       string
	Password string
}

type UserStorage interface {
	ValidateUser (ctx context.Context, userID, pwd string) (bool, error)
}