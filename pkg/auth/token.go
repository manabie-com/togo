package auth

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/manabie-com/togo/pkg/errorx"
)

const (
	AuthorizationHeader  = "Authorization"
	AuthorizationScheme  = "Bearer"
	TokenExpiresDuration = 60
)

type TokenDetails struct {
	AccessToken string
	AtExpires   int64
}

type SessionInfo struct {
	UserID int
}

type JWTCustomClaims struct {
	UserID int `json:"user_id"`
	jwt.StandardClaims
}

func GenerateToken(userID int) (*TokenDetails, error) {
	var err error
	//Creating Access Token
	atExpires := time.Now().Add(time.Minute * TokenExpiresDuration).Unix()
	atClaims := &JWTCustomClaims{
		UserID: userID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: atExpires,
		},
	}
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	accessToken, err := at.SignedString([]byte(os.Getenv("ACCESS_SECRET")))
	if err != nil {
		return nil, err
	}

	return &TokenDetails{
		AccessToken: accessToken,
		AtExpires:   atExpires,
	}, nil
}

// Get token string from Authorization Header (Ex: Bearer bncdyfcg812h3ndsya8dg68sd12...)
func getTokenStrFromHeader(header http.Header) (string, error) {
	s := header.Get(AuthorizationHeader)
	if s == "" {

		return "", errorx.ErrAuthFailure(errors.New(fmt.Sprintf("Missing authorization string.")))
	}
	splits := strings.SplitN(s, " ", 2)
	if len(splits) < 2 {
		return "", errorx.ErrAuthFailure(errors.New(fmt.Sprintf("Bad authorization string.")))
	}
	if splits[0] != AuthorizationScheme {
		return "", errorx.ErrAuthFailure(errors.New(fmt.Sprintf("Request unauthenticated with %v", AuthorizationScheme)))
	}
	return splits[1], nil
}

func getTokenFromRequest(r *http.Request) (*jwt.Token, error) {
	tokenStr, err := getTokenStrFromHeader(r.Header)
	if err != nil {
		return nil, err
	}
	atClaims := &JWTCustomClaims{}
	token, err := jwt.ParseWithClaims(tokenStr, atClaims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("ACCESS_SECRET")), nil
	})
	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				errorx.ErrAuthFailure(errors.New(fmt.Sprintf("Validation Error Malformed")))
				return nil, errorx.ErrAuthFailure(errors.New(fmt.Sprintf("Validation Error Malformed")))
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				// Token is expired

				return nil, errorx.ErrAccessTokenExpired(errors.New(fmt.Sprintf("Token have already expried")))

			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, errorx.ErrInvalidAccessToken(errors.New(fmt.Sprintf("Validation Error Not Valid Yet")))
			} else {
				return nil, errorx.ErrInvalidAccessToken(errors.New(fmt.Sprintf("Token Invalid")))
			}
		}
	}
	if time.Now().Unix() > atClaims.ExpiresAt {
		return nil, errorx.ErrInvalidAccessToken(errors.New(fmt.Sprintf("Token Invalid")))
	}
	if err != nil {
		return nil, err
	}
	return token, nil
}

func GetCustomClaimsFromRequest(r *http.Request) (*JWTCustomClaims, error) {
	token, err := getTokenFromRequest(r)
	if err != nil {
		return nil, errorx.ErrInvalidAccessToken(errors.New(fmt.Sprintf("Token Invalid")))

	}

	claim, ok := token.Claims.(*JWTCustomClaims)
	if ok && token.Valid {
		return claim, nil
	}
	return nil, nil
}

func ValidateToken(r *http.Request) error {
	token, err := getTokenFromRequest(r)
	if err != nil {
		return err
	}
	if _, ok := token.Claims.(JWTCustomClaims); !ok || !token.Valid {
		return err
	}
	return nil
}
