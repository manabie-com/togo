package request

import (
	"context"
	"net/http"

	"github.com/manabie-com/togo/internal/consts"
)

func SetUserID(req *http.Request, userId string) *http.Request {
	return req.WithContext(context.WithValue(req.Context(), consts.InjectUserID, userId))
}

func GetUserID(req *http.Request) (string, bool) {
	ctx := req.Context()
	v := ctx.Value(consts.InjectUserID)
	id, ok := v.(string)
	return id, ok
}

// Retrieve jwt secret key from request context
func SetJWTSecret(req *http.Request, jwtSecret string) *http.Request {
	return req.WithContext(context.WithValue(req.Context(), consts.InjectJWTSecret, jwtSecret))
}

// Retrieve jwt secret key from request context
func GetJWTSecret(req *http.Request) string {
	ctx := req.Context()
	v := ctx.Value(consts.InjectJWTSecret)
	value, ok := v.(string)
	if !ok {
		return ""
	}
	return value
}
