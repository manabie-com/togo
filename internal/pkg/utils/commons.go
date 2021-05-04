package pkg

import (
	"context"
	"database/sql"
	"github.com/dgrijalva/jwt-go"
	"github.com/manabie-com/togo/internal/pkg/constant"

	"log"
	"net/http"
	"time"
)

type Utils struct {
}

func (u *Utils) Value(req *http.Request, p string) sql.NullString {
	return sql.NullString{
		String: req.FormValue(p),
		Valid:  true,
	}
}

func (u *Utils) CreateToken(id string) (string, error) {
	atClaims := jwt.MapClaims{}
	atClaims["user_id"] = id
	atClaims["exp"] = time.Now().Add(time.Minute * 15).Unix()
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	token, err := at.SignedString([]byte(constant.JWTKey))
	if err != nil {
		return "", err
	}
	return token, nil
}

func ParseToken(token string) (jwt.MapClaims, error) {
	claims := make(jwt.MapClaims)
	tokenClaims, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(constant.JWTKey), nil
	})

	if tokenClaims != nil {
		if tokenClaims.Valid {
			return claims, err
		}
	}

	return nil, err
}

func (u *Utils) ValidToken(req *http.Request) (*http.Request, bool) {
	token := req.Header.Get("Authorization")

	claims := make(jwt.MapClaims)
	t, err := jwt.ParseWithClaims(token, claims, func(*jwt.Token) (interface{}, error) {
		return []byte(constant.JWTKey), nil
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

func (u *Utils) UserIDFromCtx(ctx context.Context) (string, bool) {
	v := ctx.Value("user_id")
	id, ok := v.(string)
	return id, ok
}
