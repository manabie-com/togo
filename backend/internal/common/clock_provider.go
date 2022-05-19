package common

import "go.uber.org/dig"

func ProvideClockSim(iContainer *dig.Container) error {
	return iContainer.Provide(MakeClockSim, dig.As(new(ClockI)))
}