package registry

import (
	"togo/interface/controllers"
	"togo/usecase"
)

func (r *registry) NewTodoController() controllers.TodoControllerInterface {
	return controllers.NewTodoController(r.NewTodoUsecase())
}

func (r *registry) NewTodoUsecase() usecase.TodoUsecaseInterface {
	return usecase.NewTodoUsecase(r.db)
}
