package infra

import (
	"github.com/looplab/eventhorizon"
	"github.com/looplab/eventhorizon/aggregatestore/events"
)

func ProvideAggregateStore(eventstore eventhorizon.EventStore, eventbus eventhorizon.EventBus) (eventhorizon.AggregateStore, error) {
	aggregateStore, err := events.NewAggregateStore(eventstore, eventbus)
	if err != nil {
		return nil, err
	}

	return aggregateStore, nil
}
