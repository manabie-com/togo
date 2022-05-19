package repositories

import "go.uber.org/dig"

func ProvideTaskRepository(iContainer *dig.Container) error {
	return iContainer.Provide(MakeTaskRepositorySql, dig.As(new(TaskRepositoryI)))
}