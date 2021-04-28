package validate

import (
	"context"
	"errors"
	"log"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/manabie-com/togo/define"
	"github.com/manabie-com/togo/libs"
)

func CreateToken(id string) (string, error) {
	atClaims := jwt.MapClaims{}
	atClaims["user_id"] = id
	atClaims["exp"] = time.Now().Add(time.Minute * 15).Unix()
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	token, err := at.SignedString([]byte(define.JWTKey))
	if err != nil {
		return "", err
	}
	return token, nil
}

func ValidToken(req *http.Request) (*http.Request, error) {
	token := req.Header.Get("Authorization")

	claims := make(jwt.MapClaims)
	t, err := jwt.ParseWithClaims(token, claims, func(*jwt.Token) (interface{}, error) {
		return []byte(define.JWTKey), nil
	})
	if err != nil {
		log.Println(err)
		return req, err
	}

	if !t.Valid {
		return req, errors.New("Validate token is failed")
	}

	id, ok := claims["user_id"].(string)
	if !ok {
		return req, errors.New("Validate token is failed")
	}

	req = req.WithContext(context.WithValue(req.Context(), libs.UserAuthKey(0), id))
	return req, nil
}
