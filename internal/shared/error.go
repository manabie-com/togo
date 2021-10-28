package shared

import "fmt"

type LimitedError struct {
	LimitedNumber int32
}

func (p *LimitedError) Error() string {
	return fmt.Sprintf("you are limited to create %d tasks per day", p.LimitedNumber)
}
