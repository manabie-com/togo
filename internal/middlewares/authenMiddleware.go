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
func Authen(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		valid, err := services.ValidateToken(rw, r)

		if !valid {
			log.Print(err)
			return
		}
		next.ServeHTTP(rw, r)
	})
}

// LogRequest log request info
func LogRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s %s\n", r.RemoteAddr, r.Method, r.URL)
		next.ServeHTTP(w, r)
	})
}
