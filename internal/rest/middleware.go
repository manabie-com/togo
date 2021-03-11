package rest

import (
	"context"
	"log"
	"net/http"

	"github.com/manabie-com/togo/internal/services"
)

// MethodMiddleware is a middleware designed to elegantly handle request methods
func MethodMiddleware(next http.Handler, method string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != method {
			returnErrorJSONResponse(w, "Method now allowed", http.StatusMethodNotAllowed)
			return
		}
		next.ServeHTTP(w, r)
	})
}

// AuthMiddleware acts as a middleware that requires request received from user
// must contains Authorization header, which is in form of either: `Bearer token` or `token`
func AuthMiddleware(next http.Handler, authService services.AuthService) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		returnErr := func(w http.ResponseWriter) {
			returnErrorJSONResponse(w, "Unauthorized", http.StatusUnauthorized)
		}

		token, err := normalizeAuthorizationHeader(r.Header.Get("Authorization"))
		if err != nil {
			returnErr(w)
			return
		}

		userID, err := authService.DecodeToken(r.Context(), token)
		if err != nil {
			returnErr(w)
			return
		}
		ctx := context.WithValue(r.Context(), ContextUserKey, userID)
		r = r.WithContext(ctx)
		next.ServeHTTP(w, r)
	})
}

// LoggingMiddleWare log incomming request and response
func LoggingMiddleWare(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Print(r.Method, " ", r.URL.Path)
		next.ServeHTTP(w, r)
	})
}
