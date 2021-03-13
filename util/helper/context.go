package helper

import (
	"context"

	"github.com/dgrijalva/jwt-go"
	kitjwt "github.com/go-kit/kit/auth/jwt"
)

func UserIDFromCtx(ctx context.Context) (string, bool) {
	v := ctx.Value(kitjwt.JWTClaimsContextKey)

	claims, ok := v.(jwt.MapClaims)
	if !ok {
		return "", false
	}

	id, ok := claims["user_id"].(string)
	if !ok {
		return "", false
	}

	return id, ok
}
