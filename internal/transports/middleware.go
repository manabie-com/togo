package transports

import (
	"net/http"

	"github.com/manabie-com/togo/internal/utils"
)

func Middleware(next http.Handler, jwtKey string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var ok bool
		r, ok = utils.NewJwtUtil(jwtKey).ValidToken(r)
		if !ok {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})
}
