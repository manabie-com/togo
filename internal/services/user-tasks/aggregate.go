package user_tasks

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/looplab/eventhorizon"
	"github.com/looplab/eventhorizon/aggregatestore/events"
)

type task []interface{}

const (
	DATE_LAYOUT = "2006-01-02" // YY-MM-DD
)

func init() {
	eventhorizon.RegisterAggregate(func(id uuid.UUID) eventhorizon.Aggregate {
		return &Aggregate{
			AggregateBase: events.NewAggregateBase(AggregateType, id),
			id:            uuid.UUID{},
			tasks:         make(map[string]task),
		}
	})
}

const AggregateType = eventhorizon.AggregateType("user_task")

type Aggregate struct {
	*events.AggregateBase
	id    uuid.UUID
	tasks map[string]task
}

func (a *Aggregate) HandleCommand(ctx context.Context, cmd eventhorizon.Command) error {
	switch cmd := cmd.(type) {
	case *CreateTask:
		currentDate := time.Now().Format(DATE_LAYOUT)
		if taskMap, ok := a.tasks[currentDate]; ok {
			if len(taskMap) >= cmd.TaskLimit && cmd.TaskLimit != -1 {
				return fmt.Errorf("You had %d tasks, reach limit tasks per day is %d", len(taskMap), cmd.TaskLimit)
			}
		}

		a.AppendEvent(TaskCreated, &TaskCreatedData{
			UserID:  cmd.UserID,
			Content: cmd.Content,
		}, time.Now())

	default:
		return fmt.Errorf("could not handle command: %s", cmd.CommandType())
	}

	return nil
}

func (a *Aggregate) ApplyEvent(ctx context.Context, event eventhorizon.Event) error {
	switch event.EventType() {
	case TaskCreated:
		if data, ok := event.Data().(*TaskCreatedData); ok {
			if tasks, ok := a.tasks[event.Timestamp().Format(DATE_LAYOUT)]; ok {
				tasks = append(tasks, data.Content)
				a.tasks[event.Timestamp().Format(DATE_LAYOUT)] = tasks
				return nil
			}

			a.tasks[event.Timestamp().Format(DATE_LAYOUT)] = []interface{}{data.Content}
			return nil
		} else {

		}

	default:
		return fmt.Errorf("could not apply event: %s", event.EventType())

	}

	return fmt.Errorf("invalid event data: event_type (%s), aggregate_type (%s), aggregate_id (%s)",
		event.EventType(), event.AggregateType().String(), event.AggregateID().String())
}
