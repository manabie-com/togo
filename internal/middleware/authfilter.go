package middleware

import (
	"context"
	"log"
	"net/http"

	"github.com/dgrijalva/jwt-go"
)

type AuthHandler struct {
	JWTKey  string
	handler HandlerFunc
}

func NewAuthHandler(jwtKey string, wrapper HandlerFunc) *AuthHandler {
	return &AuthHandler{JWTKey: jwtKey, handler: wrapper}
}

func (af *AuthHandler) DoFilter(resp http.ResponseWriter, req *http.Request) {
	resp.Header().Set("Access-Control-Allow-Origin", "*")
	resp.Header().Set("Access-Control-Allow-Headers", "*")
	resp.Header().Set("Access-Control-Allow-Methods", "*")

	if req.Method == http.MethodOptions {
		resp.WriteHeader(http.StatusOK)
		return
	}

	// log.Fatal("filtering authenticated token...")
	req, ok := af.validateToken(req)

	if !ok {
		log.Fatal("Unfortunately, the token is not valid")
		resp.WriteHeader(http.StatusUnauthorized)
		return
	}

	af.handler(resp, req)
}

func (af *AuthHandler) validateToken(req *http.Request) (*http.Request, bool) {
	token := req.Header.Get("Authorization")

	claims := make(jwt.MapClaims)
	t, err := jwt.ParseWithClaims(token, claims, func(*jwt.Token) (interface{}, error) {
		return []byte(af.JWTKey), nil
	})
	if err != nil {
		log.Println(err)
		return req, false
	}

	if !t.Valid {
		return req, false
	}

	id, ok := claims["user_id"].(string)
	if !ok {
		return req, false
	}

	req = req.WithContext(context.WithValue(req.Context(), UserAuthKey(0), id))
	return req, true
}

func UserIdFromCtx(ctx context.Context) (string, bool) {
	v := ctx.Value(UserAuthKey(0))
	id, ok := v.(string)
	return id, ok
}
