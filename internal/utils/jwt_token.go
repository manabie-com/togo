package utils

import (
	"context"
	"github.com/dgrijalva/jwt-go"
	"log"
	"net/http"
	"time"
)
type Jwt struct{
	jwtKey string
}

func NewJwt() *Jwt{
	return &Jwt{"AWSkey"}
}

func (j *Jwt) CreateToken(id string) (string, error) {
	atClaims := jwt.MapClaims{}
	atClaims["user_id"] = id
	atClaims["exp"] = time.Now().Add(time.Minute * 15).Unix()
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	token, err := at.SignedString([]byte(j.jwtKey))
	if err != nil {
		return "", err
	}
	return token, nil
}

func (j *Jwt) ValidToken(req *http.Request) (*http.Request, bool) {
	token := req.Header.Get("Authorization")

	claims := make(jwt.MapClaims)
	t, err := jwt.ParseWithClaims(token, claims, func(*jwt.Token) (interface{}, error) {
		return []byte(j.jwtKey), nil
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

	req = req.WithContext(context.WithValue(req.Context(), userAuthKey(0), id))
	return req, true
}
type userAuthKey int8

func UerIDFromCtx(ctx context.Context) (string, bool) {
	v := ctx.Value(userAuthKey(0))
	id, ok := v.(string)
	return id, ok
}
