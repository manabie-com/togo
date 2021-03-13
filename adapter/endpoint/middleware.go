package endpoint

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/go-kit/kit/auth/jwt"
	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"

	"github.com/valonekowd/togo/infrastructure/auth"
	"github.com/valonekowd/togo/util/helper"
)

func LoggingMiddleware(logger log.Logger) endpoint.Middleware {
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, req interface{}) (resp interface{}, err error) {
			defer func(begin time.Time) {
				logger.Log("transport_error", err, "took", time.Since(begin))
			}(time.Now().UTC())

			return next(ctx, req)
		}
	}
}

func JWTAuthMiddleware(authCfg auth.Config) endpoint.Middleware {
	m := jwt.NewParser(authCfg.JWT.Keyfunc, authCfg.JWT.SigningMethod, jwt.MapClaimsFactory)

	return func(next endpoint.Endpoint) endpoint.Endpoint {
		e := m(func(ctx context.Context, req interface{}) (interface{}, error) {
			if _, ok := helper.UserIDFromCtx(ctx); !ok {
				return nil, fmt.Errorf("checking token: %w", auth.ErrUnknownUser)
			}
			return next(ctx, req)
		})

		return func(ctx context.Context, req interface{}) (interface{}, error) {
			resp, err := e(ctx, req)
			if errors.Is(err, jwt.ErrTokenContextMissing) {
				return nil, fmt.Errorf("authenticating jwt: %w", auth.ErrNoAccessToken)
			}
			return resp, err
		}
	}
}
