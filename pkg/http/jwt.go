package http

import (
	"context"
	"net/http"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/sirupsen/logrus"
)

const UserIDKey = "userId"

func VerifyTokenAccessContext(ctx context.Context, req *http.Request) context.Context {

	claims := make(jwt.MapClaims)
	t, err := jwt.ParseWithClaims(req.Header.Get("Authorization"), claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_KEY")), nil
	})
	if err != nil || !t.Valid {
		logrus.Error(err)
		return ctx
	}

	id, ok := claims[UserIDKey].(string)
	if !ok {
		logrus.Error("userId not found")
		return ctx
	}
	return context.WithValue(ctx, UserIDKey, id)
}

func CreateToken(id string) (string, error) {
	atClaims := jwt.MapClaims{}
	atClaims[UserIDKey] = id
	atClaims["exp"] = time.Now().Add(time.Minute * 15).Unix()
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	token, err := at.SignedString([]byte(os.Getenv("JWT_KEY")))
	if err != nil {
		return "", err
	}
	return token, nil
}
