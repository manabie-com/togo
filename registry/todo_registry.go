package registry

import (
	"togo/interface/controllers"
	"togo/usecase"
)

func (r *registry) NewTodoController() controllers.TodoControllerInterface {
	return controllers.NewTodoController(r.NewTodoUsecase(), r.NewTodoLimitUsecase())
}

func (r *registry) NewTodoUsecase() usecase.TodoUsecaseInterface {
	return usecase.NewTodoUsecase(r.db)
}

func (r *registry) NewTodoLimitUsecase() usecase.TodoLimitUsecaseInterface {
	return usecase.NewTodoLimitUsecase(r.db)
}
