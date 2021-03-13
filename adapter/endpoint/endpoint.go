package endpoint

import (
	"context"

	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"

	"github.com/valonekowd/togo/infrastructure/auth"
	"github.com/valonekowd/togo/usecase"
)

type ServerEndpoint struct {
	HealthCheck endpoint.Endpoint
	User        UserEndpoint
	Task        TaskEndpoint
}

func MakeServerEndpoint(u usecase.Usecase, authCfg auth.Config, logger log.Logger) ServerEndpoint {
	return ServerEndpoint{
		HealthCheck: MakeHealthCheckEndpoint(),
		User:        NewUserEndpoint(u, authCfg, log.With(logger, "resource", "user")),
		Task:        NewTaskEndpoint(u, authCfg, log.With(logger, "resource", "task")),
	}
}

func MakeHealthCheckEndpoint() endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		return map[string]string{
			"name":   "Todo App",
			"status": "SERVING",
		}, nil
	}
}
