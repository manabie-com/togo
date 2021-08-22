package models

import (
	"github.com/google/uuid"
)

// User reflects users data from DB
type User struct {
	ID       uuid.UUID
	Name 	 string
	Email 	 string
	Password string
	AccessToken string
}
