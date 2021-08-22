package tools

import (
	"github.com/dgrijalva/jwt-go"
	"net/http"
	"time"
)

func CreateToken(id, JWTKey string) (string, *TodoError) {
	atClaims := jwt.MapClaims{}
	atClaims["user_id"] = id
	atClaims["exp"] = time.Now().Add(time.Minute * 15).Unix()
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	token, err := at.SignedString([]byte(JWTKey))
	if err != nil {
		return "", NewTodoError(http.StatusInternalServerError, err.Error())
	}
	return token, nil
}
