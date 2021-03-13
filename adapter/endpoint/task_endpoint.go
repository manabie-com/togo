package endpoint

import (
	"context"

	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"

	"github.com/valonekowd/togo/infrastructure/auth"
	"github.com/valonekowd/togo/usecase"
	"github.com/valonekowd/togo/usecase/request"
)

type TaskEndpoint struct {
	Get    endpoint.Endpoint
	Create endpoint.Endpoint
}

func NewTaskEndpoint(u usecase.Usecase, authCfg auth.Config, logger log.Logger) TaskEndpoint {
	var getEndpoint endpoint.Endpoint
	{
		getEndpoint = MakeGetTasksEndpoint(u)
		getEndpoint = JWTAuthMiddleware(authCfg)(getEndpoint)
		getEndpoint = LoggingMiddleware(log.With(logger, "method", "Get"))(getEndpoint)
	}

	var createEndpoint endpoint.Endpoint
	{
		createEndpoint = MakeCreateTaskEndpoint(u)
		createEndpoint = JWTAuthMiddleware(authCfg)(createEndpoint)
		createEndpoint = LoggingMiddleware(log.With(logger, "method", "Create"))(createEndpoint)
	}

	return TaskEndpoint{
		Get:    getEndpoint,
		Create: createEndpoint,
	}
}

func MakeGetTasksEndpoint(u usecase.Usecase) endpoint.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		r := req.(request.GetTasks)

		return u.Task.Fetch(ctx, r)
	}
}

func MakeCreateTaskEndpoint(u usecase.Usecase) endpoint.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		r := req.(request.CreateTask)

		return u.Task.Create(ctx, r)
	}
}
