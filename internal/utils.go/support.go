package utils

import "context"

type UserAuthKey int8

func GetUserIDFromCtx(ctx context.Context) (string, bool) {
	v := ctx.Value(UserAuthKey(0))
	id, ok := v.(string)
	return id, ok
}
