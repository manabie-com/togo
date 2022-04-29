package controllers

import "go.uber.org/dig"

func ProvideTaskController(iContainer *dig.Container) error {
	return iContainer.Provide(MakeTaskController)
}