package repositories

import "go.uber.org/dig"

func ProvideUserRepository(iContainer *dig.Container) error {
	return iContainer.Provide(MakeUserRepositorySql, dig.As(new(UserRepositoryI)))
}