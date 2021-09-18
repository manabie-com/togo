package api

import (
	"context"
	"net/http"
)

func RouteHandle(
	ctx context.Context,
	httpRequest *http.Request,
	decode func(context.Context, *http.Request) (interface{}, error),
	endpointHandler func(ctx context.Context, req interface{}) (interface{}, error),
) (interface{}, error) {
	req, err := decode(ctx, httpRequest)
	if err != nil {
		return nil, err
	}

	return endpointHandler(ctx, req)
}
