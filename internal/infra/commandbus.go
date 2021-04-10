package infra

import (
	"github.com/looplab/eventhorizon"
	"github.com/looplab/eventhorizon/commandhandler/bus"
	user_tasks "github.com/manabie-com/togo/internal/services/user-tasks"
)

func ProvideCommandbus(
	userTaskCommandHandler user_tasks.UserTaskCommandHandler,
) eventhorizon.CommandHandler {
	bus := NewCommandbus()

	bus.SetHandler(userTaskCommandHandler, user_tasks.CreateTaskCommand)

	return bus
}

func NewCommandbus() *bus.CommandHandler {
	return bus.NewCommandHandler()
}
