package views

import "go.uber.org/dig"

func ProvideTaskViewRest(iContainer *dig.Container) error {
	return iContainer.Provide(MakeTaskViewRest)
}