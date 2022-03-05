package jwt

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/pkg/errors"
)

const MyUserClaim = "MyUserClaim"

type DataClaim struct {
	UserID int64 `json:"user_id"`
	jwt.StandardClaims
}

func EncodeToken(userID int64, secretKey string) (string, error) {
	claims := DataClaim{
		userID,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(12 * time.Hour).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func DecodeToken(token string, secretKey string) (*DataClaim, error) {
	tokenType, err := jwt.ParseWithClaims(token, &DataClaim{}, func(_ *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})
	if err != nil {
		return nil, err
	}

	claims, ok := tokenType.Claims.(*DataClaim)
	if !ok || !tokenType.Valid {
		return nil, errors.New("failed to decode token")
	}

	return claims, nil
}
