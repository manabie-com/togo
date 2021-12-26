package utils

import "context"

type userAuthKey int8

func AddToContext(ctx context.Context, id string) context.Context {
	return context.WithValue(ctx, userAuthKey(0), id)
}

func ExtractFromContext(ctx context.Context) (string, bool) {
	v := ctx.Value(userAuthKey(0))
	id, ok := v.(string)
	return id, ok
}
