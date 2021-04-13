package usecase

import (
	"context"
	"database/sql"
	"github.com/dgrijalva/jwt-go"
	"log"
	"time"
)

func (u *uc) Validate(ctx context.Context, user, password sql.NullString) bool {
	if user.String == "" || password.String == "" {
		log.Println("invalid user name - password")
		return false
	}
	return u.task.ValidateUser(ctx, user, password)
}

func (u *uc) CreateToken(id, jwtKey string) (string, error) {
	atClaims := jwt.MapClaims{}
	atClaims["user_id"] = id
	atClaims["exp"] = time.Now().Add(time.Minute * 15).Unix()
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	token, err := at.SignedString([]byte(jwtKey))
	if err != nil {
		return "", err
	}
	return token, nil
}

func (u *uc) ValidToken(token, JWTKey string) (string, bool) {
	claims := make(jwt.MapClaims)
	t, err := jwt.ParseWithClaims(token, claims, func(*jwt.Token) (interface{}, error) {
		return []byte(JWTKey), nil
	})
	if err != nil {
		log.Println("ValidToken err: ", err)
		return "", false
	}

	if !t.Valid {
		return "", false
	}

	id, ok := claims["user_id"].(string)
	if !ok {
		return "", false
	}
	return id, true
}
