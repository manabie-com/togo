package infra

import (
	"github.com/looplab/eventhorizon"
	"github.com/manabie-com/togo/internal/services/tasks"
)

type TaskHandler interface {
	eventhorizon.EventHandler
}

func ProvideTaskHandler(taskRepo tasks.TaskRepo) TaskHandler {
	return tasks.NewTaskHandler(taskRepo)
}
