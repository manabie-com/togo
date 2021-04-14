package infra

import (
	"github.com/looplab/eventhorizon"
	"github.com/looplab/eventhorizon/eventhandler/projector"
	user_tasks "github.com/manabie-com/togo/internal/services/user-tasks"
	"github.com/manabie-com/togo/internal/services/users"
)

type UserTaskProjector interface {
	eventhorizon.EventHandler
}

func ProvideUserTaskProjector(
	userConfigRepo users.UserConfigRepo,
	userTaskRepo user_tasks.UserTaskRepo,
) UserTaskProjector {
	proj := projector.NewEventHandler(user_tasks.NewUserTaskProjector(userConfigRepo), userTaskRepo)
	proj.SetEntityFactory(func() eventhorizon.Entity {
		return &user_tasks.UserTask{}
	})

	return proj
}
