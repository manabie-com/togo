//go:build wireinject
// +build wireinject

package registry

import (
	"context"
	"github.com/google/wire"
	"net/http"
	"togo/domain/handler"
	"togo/domain/middlewares"
	"togo/domain/repositories"
	"togo/domain/services"
)

func InitHTTPServer(ctx context.Context) (http.Handler, error) {
	wire.Build(
		handler.NewHTTPServer,
		middlewares.NewMiddleware,

		services.NewTodoService,

		repositories.NewTodoRepository,
	)
	return nil, nil
}
