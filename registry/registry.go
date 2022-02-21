package registry

import (
	"togo/infrastructure/database"
	"togo/interface/controllers"
)

type registry struct {
	db database.DbInterface
}

func NewRegistry(db database.DbInterface) RegistryInterface {
	return &registry{db}
}

func (r *registry) NewAppController() controllers.AppController {
	return r.NewTodoController()
}
