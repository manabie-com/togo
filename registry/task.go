package registry

import (
	"github.com/manabie-com/togo/internal/task/repository"
	"github.com/manabie-com/togo/internal/task/service"
	repository2 "github.com/manabie-com/togo/internal/user/repository"
)

func (r *Registry) RegisterTaskService() service.TaskService {
	return service.NewTaskService(repository2.NewUserRepository(r.DB.ManabieDB), repository.NewTaskRepository(r.DB.ManabieDB), r.DB.ManabieDB)
}
