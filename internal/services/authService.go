package services

import (
	"context"
	"log"
	"net/http"

	jwt "github.com/dgrijalva/jwt-go"
)

// ValidateToken validate token
func ValidateToken(resp http.ResponseWriter, req *http.Request) {
	token := req.Header.Get("Authorization")

	claims := make(jwt.MapClaims)

	t, err := jwt.ParseWithClaims(token, claims, func(*jwt.Token) (interface{}, error) {

		return []byte(""), nil

	})

	if err != nil {
		log.Println(err)
		resp.WriteHeader(http.StatusUnauthorized)
		return
	}

	if !t.Valid {
		resp.WriteHeader(http.StatusUnauthorized)
		return
	}

	id, ok := claims["user_id"].(string)

	if !ok {
		resp.WriteHeader(http.StatusUnauthorized)
		return
	}

	req.WithContext(context.WithValue(req.Context(), userAuthKey(0), id))
}

type userAuthKey int8

func userIDFromCtx(ctx context.Context) (string, bool) {
	v := ctx.Value(userAuthKey(0))
	id, ok := v.(string)
	return id, ok
}
