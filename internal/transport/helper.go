package transport

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

const JWTKey = "wqGyEBBfPK9w3Lxw"

func CreateToken(id string) (string, error) {
	atClaims := jwt.MapClaims{}
	atClaims["user_id"] = id
	atClaims["exp"] = time.Now().Add(time.Minute * 15).Unix()
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	token, err := at.SignedString([]byte(JWTKey))
	if err != nil {
		return "", err
	}
	return token, nil
}
