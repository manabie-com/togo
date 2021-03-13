package endpoint

import (
	"context"

	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"

	"github.com/valonekowd/togo/infrastructure/auth"
	"github.com/valonekowd/togo/usecase"
	"github.com/valonekowd/togo/usecase/request"
)

type UserEndpoint struct {
	SignIn endpoint.Endpoint
	SignUp endpoint.Endpoint
}

func NewUserEndpoint(u usecase.Usecase, authCfg auth.Config, logger log.Logger) UserEndpoint {
	var signInEndpoint endpoint.Endpoint
	{
		signInEndpoint = MakeSignInEndpoint(u)
		// getTransactionsEndpoint = JWTAuthMiddleware(authCfg)(getTransactionsEndpoint)
		// getTransactionsEndpoint = LoggingMiddleware(logger)(getTransactionsEndpoint)
	}

	var signUpEndpoint endpoint.Endpoint
	{
		signUpEndpoint = MakeSignUpEndpoint(u)
		// createTransactionEndpoint = JWTAuthMiddleware(authCfg)(createTransactionEndpoint)
		// createTransactionEndpoint = LoggingMiddleware(logger)(createTransactionEndpoint)
	}

	return UserEndpoint{
		SignIn: signInEndpoint,
		SignUp: signUpEndpoint,
	}
}

func MakeSignInEndpoint(u usecase.Usecase) endpoint.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		r := req.(request.SignIn)

		return u.User.SignIn(ctx, r)
	}
}

func MakeSignUpEndpoint(u usecase.Usecase) endpoint.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		r := req.(request.SignUp)

		return u.User.SignUp(ctx, r)
	}
}
