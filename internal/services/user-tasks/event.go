package user_tasks

import (
	"github.com/google/uuid"
	"github.com/looplab/eventhorizon"
)

const (
	TaskCreated   = eventhorizon.EventType("user:task_created")
	ConfigUpdated = eventhorizon.EventType("user:config_updated")
)

func init() {
	eventhorizon.RegisterEventData(TaskCreated, func() eventhorizon.EventData {
		return &TaskCreatedData{}
	})

	eventhorizon.RegisterEventData(ConfigUpdated, func() eventhorizon.EventData {
		return &ConfigUpdatedData{}
	})
}

type TaskCreatedData struct {
	UserID  uuid.UUID `json:"user_id"`
	Content string    `json:"content"`
}

type ConfigUpdatedData struct {
	UserID    uuid.UUID `json:"user_id"`
	TaskLimit int       `json:"task_limit"`
}
