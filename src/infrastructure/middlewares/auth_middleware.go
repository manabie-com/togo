package middlewares

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/quochungphp/go-test-assignment/src/pkgs/token"
)

// AuthMiddleware ...
func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		err := token.ValidToken(r)
		if err != nil {
			w.Header().Set("Content-type", "application/json")
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode("Unauthorized")
			return
		}

		// Reflect user info to var
		tk := token.Token{}
		_, err = tk.ExtractToken(r)
		if err != nil {
			w.Header().Set("Content-type", "application/json")
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(fmt.Sprintf("Unauthorized :%s", err))
			return
		}

		next.ServeHTTP(w, r)
	})
}
