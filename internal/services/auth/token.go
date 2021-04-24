package auth

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/manabie-com/togo/internal/app/config"
	"github.com/manabie-com/togo/internal/utils"
)

var secretkey = config.LoadConfigs().Jwt.SecretKey

func CreateToken(id int) (string, error) {
	atClaims := jwt.MapClaims{}
	atClaims["user_id"] = id
	fmt.Println(id)
	atClaims["exp"] = time.Now().Add(time.Minute * 150).Unix()
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	token, err := at.SignedString([]byte(secretkey))
	if err != nil {
		return "", err
	}
	return token, nil
}

func ValidToken(req *http.Request) (*http.Request, bool) {
	token := req.Header.Get("Authorization")

	claims := make(jwt.MapClaims)
	t, err := jwt.ParseWithClaims(token, claims, func(*jwt.Token) (interface{}, error) {
		return []byte(secretkey), nil
	})
	if err != nil {
		log.Println(err)
		return req, false
	}

	if !t.Valid {
		return req, false
	}

	id, ok := claims["user_id"].(float64)

	if !ok {
		return req, false
	}

	req = req.WithContext(context.WithValue(req.Context(), utils.UserAuthKey(0), uint64(id)))
	return req, true
}
