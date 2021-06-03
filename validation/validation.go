package validation

import (
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// Claims ..
type Claims struct {
	UID string `json:"user_id,omitempty"`
	jwt.StandardClaims
}

const (
	JWTKey         = "wqGyEBBfPK9w3Lxw"
	ttlAccessToken = 30 * 24 * time.Hour
	// ErrTokenExpiredMsg ..
	ErrTokenExpiredMsg = "Token is expired"
	// ErrTokenInvalidMsg ..
	ErrTokenInvalidMsg = "Token is invalid"
)

// GenarateAccessToken ..
func GenarateAccessToken(id string) (string, error) {
	expiredTime := time.Now().Add(ttlAccessToken).Unix()
	claims := &Claims{
		UID: id,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expiredTime,
		},
	}

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := jwtToken.SignedString([]byte(JWTKey))
	if err != nil {
		return "", err
	}
	return token, nil
}

// GenerateToken ..
func GenerateToken(id string) (string, error) {
	accessToken, err := GenarateAccessToken(id)
	if err != nil {
		return "", err
	}
	return accessToken, nil
}

// VerifyToken ..
func VerifyToken(token string) (*Claims, error) {
	claims := &Claims{}
	jwtToken, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(JWTKey), nil
	})

	if err != nil {
		v := err.(*jwt.ValidationError)
		if v.Errors == jwt.ValidationErrorExpired {
			return nil, errors.New(ErrTokenExpiredMsg)
		}
		return nil, err
	}

	if !jwtToken.Valid {
		return nil, err
	}

	return claims, nil
}
