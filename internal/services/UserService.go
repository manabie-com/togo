package services

import (
	"time"
	"context"
	"net/http"
	"log"

	jwt "github.com/dgrijalva/jwt-go"	
	"github.com/manabie-com/togo/internal/storages"
)	

const JWTKey = "wqGyEBBfPK9w3Lxw"


type IUserService interface {
	ValidateUser(user_id, password string) bool
	CreateToken(user_id string) (string, error)
}

type UserService struct {
	storages.IUserRepo
}

func (service *UserService) ValidateUser(user_id string, password string) bool {
	result := service.ValidateUserRepo(user_id, password)
	return result
}

func (service *UserService)CreateToken(id string) (string, error) {
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

func ValidToken(req *http.Request) (*http.Request, bool) {
	token := req.Header.Get("Authorization")

	claims := make(jwt.MapClaims)
	t, err := jwt.ParseWithClaims(token, claims, func(*jwt.Token) (interface{}, error) {
		return []byte(JWTKey), nil
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

func UserIDFromCtx(ctx context.Context) (string, bool) {
	v := ctx.Value(userAuthKey(0))
	id, ok := v.(string)
	return id, ok
}