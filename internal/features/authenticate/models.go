package authenticate

import "github.com/google/uuid"

// User represents information about an individual user.
type User struct {
	ID           uuid.UUID
	Email        string
	PasswordHash []byte
}
