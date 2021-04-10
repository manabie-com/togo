package user_tasks

import (
	"context"
	"errors"
	"fmt"

	"github.com/manabie-com/togo/internal/services/users"

	"github.com/looplab/eventhorizon"
	"github.com/looplab/eventhorizon/eventhandler/projector"
)

type userTaskProjector struct {
	userConfigRepo users.UserConfigRepo
}

func NewUserTaskProjector(
	userConfigRepo users.UserConfigRepo,
) *userTaskProjector {
	return &userTaskProjector{
		userConfigRepo: userConfigRepo,
	}
}

func (p *userTaskProjector) ProjectorType() projector.Type {
	return projector.Type("user_task")
}

func (p *userTaskProjector) Project(
	ctx context.Context,
	event eventhorizon.Event,
	entity eventhorizon.Entity,
) (eventhorizon.Entity, error) {
	m, ok := entity.(*UserTask)

	if !ok {
		return nil, errors.New("model is of incorrect type")
	}

	switch event.EventType() {
	case TaskCreated:
		if data, ok := event.Data().(*TaskCreatedData); ok {
			if m.Version == 0 {
				m.ID = data.UserID
				m.CreatedAt = event.Timestamp()
			}

			m.UpdatedAt = event.Timestamp()
			m.NumOfTasks++
		}

	default:
		return m, fmt.Errorf("could not project event: %s", event.EventType())
	}

	m.Version++

	return m, nil
}
