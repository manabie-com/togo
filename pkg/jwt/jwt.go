package jwt

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/TrinhTrungDung/togo/pkg/server"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

// JWT provides a Json-Web-Token structure
type JWT struct {
	key      []byte
	duration time.Duration
	algo     jwt.SigningMethod
}

// New creates new JWT for authentication/authorization middleware
func New(algo, secret string, duration int) *JWT {
	signingMethod := jwt.GetSigningMethod(algo)
	if signingMethod == nil {
		panic("Invalid JWT Signing Method")
	}

	return &JWT{
		key:      []byte(secret),
		duration: time.Duration(duration),
		algo:     signingMethod,
	}
}

// MWFunc implements JWT echo middleware function
func (j *JWT) MWFunc() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			token, err := j.parseTokenFromHeader(c)
			if err != nil {
				c.Logger().Errorf("Error parsing token: %+v", err)
			}
			if !token.Valid {
				return server.NewHTTPError(http.StatusUnauthorized, "UNAUTHORIZED", "Your session is unauthorized or has been expired.").SetInternal(err)
			}

			claims := token.Claims.(jwt.MapClaims)
			for key, val := range claims {
				c.Set(key, val)
			}

			return next(c)
		}
	}
}

func (j *JWT) GenerateToken(claims map[string]interface{}) (string, int, error) {
	expTime := time.Now().Add(j.duration)
	claims["exp"] = expTime.Unix()

	token := jwt.NewWithClaims(j.algo, jwt.MapClaims(claims))
	tokenString, err := token.SignedString(j.key)

	return tokenString, int(expTime.Sub(time.Now()).Seconds()), err
}

// parseTokenFromHeader parses token from Authorization header
func (j *JWT) parseTokenFromHeader(c echo.Context) (*jwt.Token, error) {
	token := c.Request().Header.Get("Authorization")
	if token == "" {
		return nil, fmt.Errorf("Token Not Found")
	}
	parts := strings.SplitN(token, " ", 2)
	if !(len(parts) == 2 && strings.ToLower(parts[0]) == "bearer") {
		return nil, fmt.Errorf("Invalid Token")
	}

	return j.parseToken(parts[1])
}

// parseToken parses token from string
func (j *JWT) parseToken(input string) (*jwt.Token, error) {
	return jwt.Parse(input, func(token *jwt.Token) (interface{}, error) {
		if j.algo != token.Method {
			return nil, fmt.Errorf("Token Method Mismatched")
		}
		return j.key, nil
	})
}
