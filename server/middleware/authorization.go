package middleware

import (
	"net/http"

	"github.com/manabie-com/togo/pkg/customcontext"
	"github.com/manabie-com/togo/pkg/jwtprovider"
)

type authMiddleware struct {
	jwtProvider jwtprovider.JWTProvider
}

func NewAuthMiddleware(jwtProvider jwtprovider.JWTProvider) *authMiddleware {
	return &authMiddleware{
		jwtProvider: jwtProvider,
	}
}

func (m *authMiddleware) IsAuthorized(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		token := extractToken(r)
		userID, ok := m.verifyToken(token)
		if !ok {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		ctx := customcontext.SetUserIDToContext(r.Context(), userID)
		r = r.WithContext(ctx)
		next.ServeHTTP(w, r)
	})
}

func (m *authMiddleware) verifyToken(token string) (string, bool) {

	payload, valid := m.jwtProvider.Parse(token)
	if !valid {
		return "", false
	}

	userID, ok := payload["user_id"].(string)
	if !ok {
		return "", false
	}
	return userID, true
}

func extractToken(req *http.Request) string {
	return req.Header.Get("Authorization")
}
