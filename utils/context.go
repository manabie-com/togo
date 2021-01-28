package utils

import (
	"context"
	"database/sql"
	"net/http"
)

type userAuthKey int8

// UserIDFromCtx get user_id from context
func UserIDFromCtx(ctx context.Context) (string, bool) {
	v := ctx.Value(userAuthKey(0))
	id, ok := v.(string)
	return id, ok
}

// Value get value from context
func Value(req *http.Request, p string) sql.NullString {
	return sql.NullString{
		String: req.FormValue(p),
		Valid:  true,
	}
}

// NullString get null string
func NullString(s string) sql.NullString {
	return sql.NullString{
		String: s,
		Valid:  true,
	}
}
