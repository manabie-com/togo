package middlewares

import (
	"context"
	"github.com/manabie-com/togo/utils"
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
		token, err := utils.VerifyToken(r)

		if err != nil {
			utils.JSON(w, http.StatusUnauthorized, map[string]string{"message": "Invalid Token"})
			return
		}

		username := utils.ExtractTokenMetadata(token)

		ctx := context.WithValue(r.Context(), "username", username)

		next.ServeHTTP(w, r.WithContext(ctx))
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
