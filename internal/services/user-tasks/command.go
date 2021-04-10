package user_tasks

import (
	"github.com/google/uuid"
	"github.com/looplab/eventhorizon"
)

const (
	CreateTaskCommand = eventhorizon.CommandType("user:create_task")
)

type CreateTask struct {
	UserID    uuid.UUID `json:"user_id"`
	Content   string    `json:"content"`
	TaskLimit int       `json:"task_limit" eh:"optional"`
}

func (c *CreateTask) AggregateType() eventhorizon.AggregateType { return AggregateType }
func (c *CreateTask) AggregateID() uuid.UUID                    { return c.UserID }
func (c *CreateTask) CommandType() eventhorizon.CommandType     { return CreateTaskCommand }
