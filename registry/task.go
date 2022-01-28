package registry

import (
	"github.com/manabie-com/togo/internal/task/repository"
	"github.com/manabie-com/togo/internal/task/service"
)

func (r *Registry) RegisterTaskService() service.TaskService {
	return service.NewTaskService(repository.NewTaskRepository(r.DB.ManabieDB), r.DB.ManabieDB)
}
