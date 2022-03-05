package util

import (
	"fmt"
	"time"

	"github.com/labstack/echo/v4"
)

type CustomEchoContext struct {
	echo.Context
}

func (c *CustomEchoContext) Deadline() (deadline time.Time, ok bool) {
	return c.Request().Context().Deadline()
}

func (c *CustomEchoContext) Done() <-chan struct{} {
	return c.Request().Context().Done()
}

func (c *CustomEchoContext) Err() error {
	return c.Request().Context().Err()
}

func (c *CustomEchoContext) Value(key interface{}) interface{} {
	return c.Get(fmt.Sprintf("%v", key))
}
