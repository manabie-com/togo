package services

import (
	"context"
	"log"
	"net/http"

	jwt "github.com/dgrijalva/jwt-go"
)

// ValidateToken validate token
func ValidateToken(resp http.ResponseWriter, req *http.Request) (bool, error) {
	token := req.Header.Get("Authorization")

	claims := make(jwt.MapClaims)

	t, err := jwt.ParseWithClaims(token, claims, func(*jwt.Token) (interface{}, error) {

		return []byte("123"), nil

	})

	if err != nil {
		log.Println(err)
		http.Error(resp, "UnAuthorization", http.StatusUnauthorized)
		return false, err
	}

	if !t.Valid {
		http.Error(resp, "UnAuthorization", http.StatusUnauthorized)
		return false, nil
	}

	id, ok := claims["userId"].(string)

	if !ok {
		http.Error(resp, "UnAuthorization", http.StatusUnauthorized)
		return false, nil
	}

	req.WithContext(context.WithValue(req.Context(), userAuthKey(0), id))

	return true, nil
}

type userAuthKey int8

func userIDFromCtx(ctx context.Context) (string, bool) {
	v := ctx.Value(userAuthKey(0))
	id, ok := v.(string)
	return id, ok
}
