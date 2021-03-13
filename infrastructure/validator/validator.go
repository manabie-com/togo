package validator

import "context"

type Validator interface {
	Struct(ctx context.Context, s interface{}) error
}
