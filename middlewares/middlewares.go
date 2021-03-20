package middlewares

import (
	"log"
	"net/http"
	"time"
)

func LoggingMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		log.Printf("Started %s %s", r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
		log.Printf("Completed %s in %v", r.URL.Path, time.Since(start))
	}
}

func AuthMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//TODO: Add authentication
		log.Println("Authentication required")
		next.ServeHTTP(w, r)
	}
}

type Middleware func(http.HandlerFunc) http.HandlerFunc

func Middlewares(next http.HandlerFunc, m ...Middleware) http.HandlerFunc {

	if len(m) < 1 {
		return next
	}

	wrapped := next

	// loop in reverse to preserve middleware order
	for i := len(m) - 1; i >= 0; i-- {
		wrapped = m[i](wrapped)
	}

	return wrapped

}
