package infra

import (
	"context"

	user_tasks "github.com/manabie-com/togo/internal/services/user-tasks"

	"github.com/looplab/eventhorizon"
	"github.com/looplab/eventhorizon/eventbus/local"
)

func ProvideEventBus(
	userTaskProjector UserTaskProjector,
	taskHandler TaskHandler,
) eventhorizon.EventBus {
	eventbus := NewInMemEventbus()

	eventbus.AddHandler(context.Background(), eventhorizon.MatchEvents{
		user_tasks.TaskCreated,
	}, userTaskProjector)

	eventbus.AddHandler(context.Background(), eventhorizon.MatchEvents{
		user_tasks.TaskCreated,
	}, taskHandler)

	return eventbus
}

func NewInMemEventbus() *local.EventBus {
	return local.NewEventBus()
}
