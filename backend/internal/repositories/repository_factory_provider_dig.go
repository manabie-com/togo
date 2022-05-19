package repositories

import "go.uber.org/dig"

func ProvideRepositoryFactory(iContainer *dig.Container) error {
	return iContainer.Provide(MakeRepositoryFactorySql, dig.As(new(RepositoryFactoryI)))
}

func ProvidMockeRepositoryFactory(iContainer *dig.Container) error {
	return iContainer.Provide(MakeRepositoryFactoryMock, dig.As(new(RepositoryFactoryI)))
}