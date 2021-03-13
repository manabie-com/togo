package decoder

import (
	"context"
	"net/http"

	kithttp "github.com/go-kit/kit/transport/http"
)

func ValidatingMiddleware(
	next kithttp.DecodeRequestFunc,
	validateFunc func(context.Context, interface{}) error,
) kithttp.DecodeRequestFunc {
	return func(ctx context.Context, req *http.Request) (interface{}, error) {
		r, err := next(ctx, req)
		if err != nil {
			return nil, err
		}

		if err = validateFunc(ctx, r); err != nil {
			return nil, err
		}

		return r, nil
	}
}
