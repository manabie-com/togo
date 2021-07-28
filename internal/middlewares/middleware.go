package middleware

import (
	"context"
	"net/http"
	"time"

	"github.com/manabie-com/togo/internal/env"
	"github.com/manabie-com/togo/internal/pkg/logging"

	"github.com/dgrijalva/jwt-go"
)

type JWTClaims struct {
	UserID string `json:"user_id"`
	jwt.StandardClaims
}

func CORS(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "*")
		w.Header().Set("Access-Control-Allow-Methods", "*")

		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}

		h.ServeHTTP(w, r)
	})
}

func Authorization(handlerFunc http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		headerAuthorization := r.Header.Get("Authorization")
		if headerAuthorization == "" {
			http.Error(w, "Missing token", http.StatusUnauthorized)
			return
		}

		token, err := decodeToken(headerAuthorization)
		if err != nil {
			logging.Errorln("JWT Authorized", err.Error())
			http.Error(w, "Invalid token", http.StatusUnauthorized)

			return
		}

		if token.Valid {
			claims, ok := token.Claims.(*JWTClaims)
			if !ok {
				return
			}

			r = r.WithContext(context.WithValue(r.Context(), 0, claims.UserID))
			handlerFunc(w, r)
		}
	}
}

func decodeToken(tokenHeader string) (*jwt.Token, error) {
	return jwt.ParseWithClaims(
		tokenHeader,
		&JWTClaims{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(env.SecretKeyJWT()), nil
		},
	)
}

// GenerateJWT the Claims
func GenerateJWT(claims JWTClaims) (map[string]string, error) {
	claims.ExpiresAt = time.Now().Add(15 * time.Minute).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	accessToken, err := token.SignedString([]byte(env.SecretKeyJWT()))
	if err != nil {
		return nil, err
	}

	refreshToken := jwt.New(jwt.SigningMethodHS256)
	rtClaims := refreshToken.Claims.(jwt.MapClaims)
	rtClaims["sub"] = 1
	rtClaims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	rfToken, err := refreshToken.SignedString([]byte(env.SecretKeyJWT()))
	if err != nil {
		return nil, err
	}

	return map[string]string{
		"access_token":  accessToken,
		"refresh_token": rfToken,
	}, nil
}
