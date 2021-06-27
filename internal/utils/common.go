package utils

import (
	"context"
	"net/http"
)

type UserAuthKey int8

func UserIDFromCtx(ctx context.Context) (string, bool) {
	v := ctx.Value(UserAuthKey(0))
	id, ok := v.(string)
	return id, ok
}

func Value(req *http.Request, p string) string {
	return req.FormValue(p)
}
