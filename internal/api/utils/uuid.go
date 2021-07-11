package utils

import "github.com/google/uuid"

func GenerateNewUUID() string {
	return uuid.New().String()
}
