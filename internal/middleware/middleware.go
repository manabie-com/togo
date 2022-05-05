package middleware

import (
	"context"
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/jmsemira/togo/internal/auth"
	"github.com/jmsemira/togo/internal/helper"
	"net/http"
	"strings"
)

func AccessTokenMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		response := map[string]interface{}{}
		// Parse JWT from extracted token string
		bearerToken, err := extractBearerToken(r)
		if err != nil {
			fmt.Println("1")
			response["status"] = "error"
			response["err_msg"] = err.Error()
			w.WriteHeader(http.StatusUnauthorized)
			helper.ReturnJSON(w, response)
			return
		}

		claims, err := verifyToken(*bearerToken)

		if err != nil {
			fmt.Println("2")
			response["status"] = "error"
			response["err_msg"] = err.Error()
			w.WriteHeader(http.StatusUnauthorized)
			helper.ReturnJSON(w, response)
			return
		}

		ctx := context.WithValue(r.Context(), "user", claims)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// Extract Bearer token from request
func extractBearerToken(r *http.Request) (*string, error) {
	authorizationHeader := r.Header.Get("authorization")

	if authorizationHeader == "" {
		return nil, errors.New("Missing Authorization header")
	}

	bearerToken := strings.Split(authorizationHeader, " ")

	if len(bearerToken) != 2 || bearerToken[0] != "Bearer" {
		return nil, errors.New("Invalid Authorization header")
	}

	return &bearerToken[1], nil
}

func verifyToken(strToken string) (*auth.Claims, error) {
	key := []byte("this is my secret key")

	claims := &auth.Claims{}

	token, err := jwt.ParseWithClaims(strToken, claims, func(token *jwt.Token) (interface{}, error) {
		return key, nil
	})
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			return claims, fmt.Errorf("invalid token signature")
		}
	}

	if !token.Valid {
		return claims, fmt.Errorf("invalid token")
	}

	return claims, nil
}
