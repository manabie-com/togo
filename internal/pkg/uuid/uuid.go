package uuid

import (
	"github.com/google/uuid"
)

// New return a string form of UUID
func New() string {
	return uuid.New().String()
}
