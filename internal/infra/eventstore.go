package infra

import (
	"github.com/go-pg/pg"
	"github.com/looplab/eventhorizon"
	"github.com/manabie-com/togo/internal/pkg/eventstore"
)

func ProvideEventStore(db *pg.DB) (eventhorizon.EventStore, error) {
	store, err := eventstore.NewEventStore(db)
	if err != nil {
		return nil, err
	}

	return store, nil
}
