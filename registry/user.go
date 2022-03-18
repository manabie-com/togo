package registry

import (
	taskrepository "github.com/manabie-com/togo/internal/task/repository"
	"github.com/manabie-com/togo/internal/user/repository"
	"github.com/manabie-com/togo/internal/user/service"
)

func (r *Registry) RegisterUserService() service.UserService {
	return service.NewUserService(repository.NewUserRepository(r.DB.ManabieDB), taskrepository.NewTaskRepository(r.DB.ManabieDB), r.DB.ManabieDB)
}
