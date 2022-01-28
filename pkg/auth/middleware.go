package auth

import (
	"context"
	"net/http"

	"github.com/manabie-com/togo/pkg/httpx"
)

func TokenAuthMiddleware(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		err := ValidateToken(r)
		if err != nil {
			httpx.WriteError(w, err)
			return
		}
		claims, err := GetCustomClaimsFromRequest(r)
		if err != nil {
			httpx.WriteError(w, err)
			return
		}
		// TODO: Get user info and pass into SessionInfo
		sessionInfo := &SessionInfo{}
		if claims != nil {
			sessionInfo.UserID = claims.UserID
		}
		ctx = context.WithValue(ctx, "ss", sessionInfo)
		next.ServeHTTP(w, r.WithContext(ctx))
	}
	return http.HandlerFunc(fn)
}
