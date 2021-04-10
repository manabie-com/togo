package util

import "github.com/looplab/eventhorizon"

func UnwrapAggError(err error) error {
	if err == nil {
		return nil
	}

	if aggErr, ok := err.(eventhorizon.AggregateError); ok {
		return aggErr.Unwrap()
	}

	return err
}
