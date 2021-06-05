package services

import (
	"context"
	"database/sql"
	"encoding/json"
	"net/http"
)

func createErrorResponse(resp http.ResponseWriter, err error) {
	resp.WriteHeader(http.StatusInternalServerError)
	json.NewEncoder(resp).Encode(map[string]string{
		"error": err.Error(),
	})
}

func value(req *http.Request, p string) sql.NullString {
	return sql.NullString{
		String: req.FormValue(p),
		Valid:  true,
	}
}

type userAuthKey int8

func userIDFromCtx(ctx context.Context) (string, bool) {
	v := ctx.Value(userAuthKey(0))
	id, ok := v.(string)
	return id, ok
}
