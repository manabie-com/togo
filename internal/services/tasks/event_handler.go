package tasks

import (
	"context"
	"errors"
	"fmt"

	user_tasks "github.com/manabie-com/togo/internal/services/user-tasks"

	"github.com/looplab/eventhorizon"
)

type taskHandler struct {
	taskRepo TaskRepo
}

func NewTaskHandler(taskRepo TaskRepo) *taskHandler {
	return &taskHandler{
		taskRepo: taskRepo,
	}
}

func (t *taskHandler) HandlerType() eventhorizon.EventHandlerType {
	return eventhorizon.EventHandlerType("task")
}

func (t *taskHandler) HandleEvent(ctx context.Context, event eventhorizon.Event) error {
	switch event.EventType() {
	case user_tasks.TaskCreated:
		if data, ok := event.Data().(*user_tasks.TaskCreatedData); ok {
			return t.taskRepo.Save(ctx, &Task{
				UserID:    data.UserID,
				Content:   data.Content,
				CreatedAt: event.Timestamp(),
			})
		} else {
			return errors.New(fmt.Sprintf("Cannot cast data %+v from event %s", event.Data(), event.EventType()))
		}
	}

	return nil
}
