package tools

import (
	"context"
	"net/http"
)

type UserAuthKey int8

func UserIDFromCtx(ctx context.Context) (string, *TodoError) {
	v := ctx.Value(UserAuthKey(0))
	id, ok := v.(string)
	if !ok {
		return "", NewTodoError(http.StatusInternalServerError, "Something wrong when set user id")
	}
	return id, nil
}

func WriteUserIDToContext(ctx context.Context, id string) context.Context {
	return context.WithValue(ctx, UserAuthKey(0), id)
}
