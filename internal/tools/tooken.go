package tools

import (
	"github.com/dgrijalva/jwt-go"
	"log"
	"net/http"
	"time"
)

type ITokenTool interface {
	CreateToken(id, JWTKey string) (string, *TodoError)
	GetToken(req *http.Request) string
	ClaimToken(token, JWTKey string) (string, *TodoError)
}

type TokenTool struct{}

func (tt *TokenTool) CreateToken(id, JWTKey string) (string, *TodoError) {
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

func (tt *TokenTool) GetToken(req *http.Request) string {
	return req.Header.Get("Authorization")
}

func (tt *TokenTool) ClaimToken(token, JWTKey string) (string, *TodoError) {
	claims := make(jwt.MapClaims)
	t, err := jwt.ParseWithClaims(token, claims, func(*jwt.Token) (interface{}, error) {
		return []byte(JWTKey), nil
	})
	if err != nil {
		log.Println(err)
		return "", NewTodoError(http.StatusInternalServerError, err.Error())
	}

	if !t.Valid {
		return "", NewTodoError(http.StatusUnauthorized, "Your request is unauthorized")
	}

	id, ok := claims["user_id"].(string)
	if !ok {
		return "", NewTodoError(http.StatusInternalServerError, "Something went wrong")
	}
	return id, nil
}

func NewTokenTool() ITokenTool {
	return &TokenTool{}
}
