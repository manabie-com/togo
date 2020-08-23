package services

import (
	"context"
	"log"
	"net/http"

	jwt "github.com/dgrijalva/jwt-go"
)

// ValidateToken validate token
func ValidateToken(resp http.ResponseWriter, req *http.Request) (bool, *http.Request) {
	token := req.Header.Get("Authorization")

	claims := make(jwt.MapClaims)

	t, err := jwt.ParseWithClaims(token, claims, func(*jwt.Token) (interface{}, error) {

		return []byte("123"), nil

	})

	if err != nil {
		log.Println(err)
		http.Error(resp, "UnAuthorization", http.StatusUnauthorized)
		return false, req
	}

	if !t.Valid {
		http.Error(resp, "UnAuthorization", http.StatusUnauthorized)
		return false, req
	}

	id, ok := claims["userId"].(string)

	if !ok {
		http.Error(resp, "UnAuthorization", http.StatusUnauthorized)
		return false, req
	}

	req = req.WithContext(context.WithValue(req.Context(), UserAuthKey("userId"), id))

	return true, req
}

// UserAuthKey struct for key which using in Context
type UserAuthKey string

// UserIDFromCtx extract userId from context
func UserIDFromCtx(ctx context.Context) (string, bool) {
	v := ctx.Value(UserAuthKey("userId"))
	id, ok := v.(string)
	return id, ok
}
