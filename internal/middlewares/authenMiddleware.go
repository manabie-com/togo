package middlewares

import (
	"log"
	"net/http"

	"github.com/manabie-com/togo/internal/services"
)

// AuthenMiddleware handles authentication
type AuthenMiddleware struct {
}

// Authen verify authentication
func (auth *AuthenMiddleware) Authen(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		services.ValidateToken(rw, r)
		next.ServeHTTP(rw, r)
	})
}

func LogRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s %s\n", r.RemoteAddr, r.Method, r.URL)
		next.ServeHTTP(w, r)
	})
}
