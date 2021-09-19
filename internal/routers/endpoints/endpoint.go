package endpoints

import "context"

type Endpoint func(ctx context.Context, req interface{}) (interface{}, error)

type Endpoints struct {
	User UserEndpoint
	Task TaskEndpoint
}
