package jwt

import (
	"errors"
	"fmt"
	"time"
	"togo/src/common/types"

	"github.com/dgrijalva/jwt-go"
	JWT "github.com/dgrijalva/jwt-go"
)

func GenerateTokenUser(userJson types.JSON) (string, error) {
	// secretKey := os.Getenv("JWT_SECRET_KEY")
	secretKey := "secrectKey"
	mySigningKey := []byte(secretKey)

	//  token is valid for 7days
	ExpiresAt := time.Now().Add(time.Hour * 24 * 7)

	claims := JWT.MapClaims{
		"user": userJson,
		"exp":  ExpiresAt.Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString(mySigningKey)
	return ss, err
}

func ValidateToken(tokenString string) (types.JSON, error) {
	secretKey := "secrectKey"
	mySigningKey := []byte(secretKey)
	token, err := JWT.Parse(tokenString, func(token *JWT.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return mySigningKey, nil
	})

	if err != nil {
		return types.JSON{}, err
	}

	if !token.Valid {
		return types.JSON{}, errors.New("invalid token")
	}

	return token.Claims.(jwt.MapClaims), nil
}
