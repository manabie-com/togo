package uuidx

import (
	"fmt"
	
	"github.com/google/uuid"
)

type UUID string

func (u UUID) String() string {
	return string(u)
}
func (u UUID) Validate() error {
	_, err := uuid.Parse(u.String())
	if err != nil {
		return fmt.Errorf("invalid uuidx: %w", err)
	}

	return nil
}
