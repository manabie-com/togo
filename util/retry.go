package util

import (
	"fmt"
	"time"
)

func Retry(attempts int, sleep time.Duration, f func() error) (err error) {
	for i := 0; ; i++ {
		err = f()
		if err == nil {
			return
		}

		if i >= (attempts - 1) {
			break
		}
		time.Sleep(sleep)
	}

	return fmt.Errorf("after %d attempts, last error: %s", attempts, err)
}

