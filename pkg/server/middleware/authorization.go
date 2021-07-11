package middleware

import (
	"context"
	"net/http"

	"github.com/manabie-com/togo/pkg/jwtprovider"
)

type userAuthKey int8

func UserIDFromCtx(ctx context.Context) (string, bool) {
	v := ctx.Value(userAuthKey(0))
	id, ok := v.(string)
	return id, ok
}

type authMiddleware struct {
	jwtProvider jwtprovider.JWTProvider
}

func NewAuthMiddleware(jwtProvider jwtprovider.JWTProvider) *authMiddleware {
	return &authMiddleware{
		jwtProvider: jwtProvider,
	}
}

func (m *authMiddleware) Authorization(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		r, ok := m.validToken(r)
		if !ok {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func (m *authMiddleware) validToken(req *http.Request) (*http.Request, bool) {
	token := req.Header.Get("Authorization")

	payload, valid := m.jwtProvider.Parse(token)
	if !valid {
		return req, false
	}

	id, ok := payload["user_id"].(string)
	if !ok {
		return req, false
	}

	req = req.WithContext(context.WithValue(req.Context(), userAuthKey(0), id))
	return req, true
}
