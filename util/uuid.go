package util

import "github.com/google/uuid"

// NewUUID - create new uuid in string
func NewUUID() string {
	return uuid.New().String()
}
