package user_tasks

import (
	"time"

	"github.com/google/uuid"
)

type UserTask struct {
	ID         uuid.UUID `json:"id"`
	Version    int       `json:"version"`
	NumOfTasks int       `json:"num_of_tasks"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

// EntityID implements the EntityID method of the eh.Entity interface.
func (e *UserTask) EntityID() uuid.UUID {
	return e.ID
}

// AggregateVersion implements the AggregateVersion method of the
// eh.Versionable interface.
func (e *UserTask) AggregateVersion() int {
	return e.Version
}
