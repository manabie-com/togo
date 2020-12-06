package middleware

import (
	"context"
	"github.com/HoangVyDuong/togo/pkg/kit"
)

func LimitCreateTask(endpoint kit.Endpoint) kit.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		return nil, nil
	}
}

