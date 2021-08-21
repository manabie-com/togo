package tools

import (
	"context"
)

type UserAuthKey int8

func UserIDFromCtx(ctx context.Context) (string, error) {
	v := ctx.Value(UserAuthKey(0))
	id, ok := v.(string)
	if !ok {
		return "", NewTodoError(500, "Something wrong when set user id")
	}
	return id, nil
}

func WriteUserIDToContext(ctx context.Context, id string) context.Context {
	return context.WithValue(ctx, UserAuthKey(0), id)
}
