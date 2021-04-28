package libs

import (
	"context"
	"database/sql"
	"net/http"
)

type UserAuthKey int8

func UserIDFromCtx(ctx context.Context) (string, bool) {

	v := ctx.Value(UserAuthKey(0))
	id, ok := v.(string)

	return id, ok
}

func Value(req *http.Request, p string) sql.NullString {
	return sql.NullString{
		String: req.FormValue(p),
		Valid:  true,
	}
}
