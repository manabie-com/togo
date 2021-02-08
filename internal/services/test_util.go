package services

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/manabie-com/togo/internal/config"
)

var ServiceMockForTest, _ = NewToDoServices(config.Jwt, config.DBType.Postgres, config.GetPostgresDBConfig().ToString())

func CreateTokenForTest(id string, exp int64) (string, error) {
	atClaims := jwt.MapClaims{}
	atClaims["user_id"] = id
	atClaims["exp"] = exp
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	token, err := at.SignedString([]byte(config.Jwt))
	if err != nil {
		return "", err
	}
	return token, nil
}
