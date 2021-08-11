package context

import (
	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/mock"
	"time"
)

type ContextMock struct {
	mock.Mock
}

func (c *ContextMock) Deadline() (deadline time.Time, ok bool) {
	panic("implement me")
}

func (c *ContextMock) Done() <-chan struct{} {
	panic("implement me")
}

func (c *ContextMock) Err() error {
	panic("implement me")
}

func (c *ContextMock) Value(key interface{}) interface{} {
	panic("implement me")
}

func (c *ContextMock) GetDb() *gorm.DB {
	panic("implement me")
}

func (c *ContextMock) GetUserId() string {
	args := c.Called()
	return args.Get(0).(string)
}
