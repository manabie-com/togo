package context

import (
	"context"
	"github.com/jinzhu/gorm"
	"github.com/manabie-com/togo/common/middleware"
)

type Context interface {
	context.Context
	GetDb() *gorm.DB
	GetUserId() string
}

func FromContext(ctx context.Context) Context {
	if wrapCtx, ok := ctx.(*contextWrapper); ok {
		return wrapCtx
	}
	return &contextWrapper{
		ctx,
	}
}

type contextWrapper struct {
	context.Context
}

func (c *contextWrapper) GetUserId() string {
	return middleware.UserIdFromAccessToken(c)
}

func (c *contextWrapper) GetDb() *gorm.DB {
	return middleware.DbFromContext(c)
}
