package infra

import (
	"github.com/looplab/eventhorizon"
	user_tasks "github.com/manabie-com/togo/internal/services/user-tasks"
)

func ProvideUserTaskCommandHandler(aggregateStore eventhorizon.AggregateStore) (user_tasks.UserTaskCommandHandler, error) {
	return user_tasks.NewUserTaskCommandHandler(aggregateStore)
}
