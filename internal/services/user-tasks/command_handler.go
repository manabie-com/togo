package user_tasks

import (
	"github.com/looplab/eventhorizon"
	"github.com/looplab/eventhorizon/commandhandler/aggregate"
)

type UserTaskCommandHandler interface {
	eventhorizon.CommandHandler
}

func NewUserTaskCommandHandler(aggregateStore eventhorizon.AggregateStore) (UserTaskCommandHandler, error) {
	handler, err := aggregate.NewCommandHandler(AggregateType, aggregateStore)
	if err != nil {
		return nil, err
	}

	return handler, nil
}
