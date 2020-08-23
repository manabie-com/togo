package middlewares

import (
	"log"
	"net/http"

	"github.com/manabie-com/togo/internal/services"
)

// Authen verify authentication
func Authen(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {

		valid, req := services.ValidateToken(rw, r)

		if !valid {
			return
		}

		next.ServeHTTP(rw, req)
	})
}

// LogRequest log request info
func LogRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		log.Printf("%s %s %s\n", r.RemoteAddr, r.Method, r.URL)
		next.ServeHTTP(w, r)
	})
}
