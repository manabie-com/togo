package http

import (
	"github.com/gin-gonic/gin"
	"github.com/shanenoi/togo/config"
	"reflect"
)

func BackgroundConfig(fn gin.HandlerFunc, configs *config.ThirdAppAdapter) gin.HandlerFunc {
	// accept specify configs
	return func(ctx *gin.Context) {
		ctx.Set(config.POSTGRESQL_DB, configs.Database)
		ctx.Set(config.REDIS_DB, configs.Redis)
		fn(ctx)
	}
}

type Route struct {
	JwtValidation bool
	HandlerFunc   gin.HandlerFunc
	MethodName    string
	RelativePath  string
}

func (r *Route) Register(root *groupRoute) {
	handler := r.HandlerFunc
	handler = BackgroundConfig(handler, root.Config)

	if r.JwtValidation {
		handler = JwtValidator(handler)
	}

	params := []reflect.Value{
		reflect.ValueOf(r.RelativePath),
		reflect.ValueOf(handler),
	}

	reflect.ValueOf(root.Group).
		MethodByName(r.MethodName).
		Call(params)
}

type groupRoute struct {
	Group  *gin.RouterGroup
	Config *config.ThirdAppAdapter
}

type RouterGroup interface {
	Load(routes ...Route)
}

func NewRouterGroup(group *gin.RouterGroup, configs *config.ThirdAppAdapter) RouterGroup {
	return &groupRoute{group, configs}
}

func (gr *groupRoute) Load(routes ...Route) {
	for _, route := range routes {
		route.Register(gr)
	}
}
