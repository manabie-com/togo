package middleware

import (
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"net/http"
	"os"
)

// LoggingMiddleware wrap the gorilla Logging Handler
func LoggingMiddleware() mux.MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			next = handlers.LoggingHandler(os.Stdout, next) // run logger before execute the logic
			next.ServeHTTP(w, r)
		})
	}
}
