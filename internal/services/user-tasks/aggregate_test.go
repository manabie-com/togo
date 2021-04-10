package user_tasks

import (
	"context"
	"testing"
	"time"

	"github.com/looplab/eventhorizon/mocks"

	"github.com/google/uuid"
	"github.com/looplab/eventhorizon"
	"github.com/looplab/eventhorizon/aggregatestore/events"
)

func TestAggregate_HandleCommand(t *testing.T) {
	timeNow := time.Now()
	id := uuid.New()
	aggBase := events.NewAggregateBase(AggregateType, id)

	type fields struct {
		AggregateBase *events.AggregateBase
		id            uuid.UUID
		tasks         map[string]task
	}
	type args struct {
		ctx context.Context
		cmd eventhorizon.Command
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "Unknow command",
			fields: fields{
				AggregateBase: aggBase,
				id:            id,
				tasks: map[string]task{
					timeNow.Format(DATE_LAYOUT): []interface{}{"hello"},
				},
			},
			args: args{
				ctx: context.Background(),
				cmd: mocks.Command{
					ID:      id,
					Content: "test",
				},
			},
			wantErr: true,
		},

		{
			name: "Not reach limit",
			fields: fields{
				AggregateBase: aggBase,
				id:            id,
				tasks: map[string]task{
					time.Now().Format(DATE_LAYOUT): []interface{}{"hello"},
				},
			},
			args: args{
				ctx: context.Background(),
				cmd: &CreateTask{
					UserID:    id,
					Content:   "hi",
					TaskLimit: 5,
				},
			},
			wantErr: false,
		},

		{
			name: "Reach limit",
			fields: fields{
				AggregateBase: aggBase,
				id:            id,
				tasks: map[string]task{
					time.Now().Format(DATE_LAYOUT): []interface{}{"hello", "hi", "ni hao ma"},
				},
			},
			args: args{
				ctx: context.Background(),
				cmd: &CreateTask{
					UserID:    id,
					Content:   "bye bye",
					TaskLimit: 3,
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &Aggregate{
				AggregateBase: tt.fields.AggregateBase,
				id:            tt.fields.id,
				tasks:         tt.fields.tasks,
			}
			if err := a.HandleCommand(tt.args.ctx, tt.args.cmd); (err != nil) != tt.wantErr {
				t.Errorf("HandleCommand() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestAggregate_ApplyEvent(t *testing.T) {
	timeNow := time.Now()
	id := uuid.New()
	aggBase := events.NewAggregateBase(AggregateType, id)

	type TaskCreatedDataNew struct {
		UserID  uuid.UUID `json:"user_id"`
		Content int       `json:"content"`
	}

	type fields struct {
		AggregateBase *events.AggregateBase
		id            uuid.UUID
		tasks         map[string]task
	}
	type args struct {
		ctx   context.Context
		event eventhorizon.Event
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "Unknow event",
			fields: fields{
				AggregateBase: aggBase,
				id:            id,
				tasks:         map[string]task{},
			},
			args: args{
				ctx: context.Background(),
				event: eventhorizon.NewEvent(eventhorizon.EventType("Unknow"), nil, timeNow,
					eventhorizon.ForAggregate(AggregateType, id, 3)),
			},
			wantErr: true,
		},

		{
			name: "happy case",
			fields: fields{
				AggregateBase: aggBase,
				id:            id,
				tasks:         map[string]task{},
			},
			args: args{
				ctx: context.Background(),
				event: eventhorizon.NewEvent(TaskCreated, &TaskCreatedData{}, timeNow,
					eventhorizon.ForAggregate(AggregateType, id, 3)),
			},
			wantErr: false,
		},

		{
			name: "unmarshal data",
			fields: fields{
				AggregateBase: aggBase,
				id:            id,
				tasks:         map[string]task{},
			},
			args: args{
				ctx: context.Background(),
				event: eventhorizon.NewEvent(TaskCreated, &TaskCreatedDataNew{}, timeNow,
					eventhorizon.ForAggregate(AggregateType, id, 3)),
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &Aggregate{
				AggregateBase: tt.fields.AggregateBase,
				id:            tt.fields.id,
				tasks:         tt.fields.tasks,
			}
			if err := a.ApplyEvent(tt.args.ctx, tt.args.event); (err != nil) != tt.wantErr {
				t.Errorf("ApplyEvent() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
