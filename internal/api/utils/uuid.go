package utils

import "github.com/google/uuid"

type GenerateNewUUIDFn func() string

func GenerateNewUUID() string {
	return uuid.New().String()
}
