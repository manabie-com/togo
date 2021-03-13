package helper

import (
	"fmt"

	_uuid "github.com/google/uuid"
)

func NewUUID() (string, error) {
	uuid, err := _uuid.NewRandom()
	if err != nil {
		return "", fmt.Errorf("generating new uuid: %w", err)
	}
	return uuid.String(), nil
}
