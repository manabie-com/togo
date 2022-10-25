package utils

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

var JwtSecret = "togo-jwt-secret"

const (
	ERROR_AUTH_CHECK_TOKEN_FAIL    = 2002
	ERROR_AUTH_CHECK_TOKEN_TIMEOUT = 2003
	ERROR_AUTH_TOKEN               = 2004
)

var MsgFlags = map[int]string{
	ERROR_AUTH_CHECK_TOKEN_FAIL:    "Token is invalid",
	ERROR_AUTH_CHECK_TOKEN_TIMEOUT: "Token is time out, please login and try again",
	ERROR_AUTH_TOKEN:               "Token is error, please try again",
}
var (
	TokenExpired     error = errors.New("Token is expired")
	TokenNotValidYet error = errors.New("Token not active yet")
	TokenMalformed   error = errors.New("That's not even a token")
	TokenInvalid     error = errors.New("Couldn't handle this token:")
)

type JWT struct {
	SigningKey []byte
}

type CustomClaims struct {
	Id       uint   `json:"id"`
	Email    string `json:"email"`
	FullName string `json:"full_name"`
	jwt.StandardClaims
}

type CustomClaimsShort struct {
	Id uint `json:"id"`
	jwt.StandardClaims
}

type TokenInfo struct {
	Token     string    `json:"token"`
	ExpiredAt time.Time `json:"expired_at"`
}

func NewJWT() *JWT {
	return &JWT{
		SigningKey: []byte(JwtSecret),
	}
}

func NewJWTSignedWithPasswordHash(hash string) *JWT {
	return &JWT{
		SigningKey: []byte(hash),
	}
}

func (j *JWT) GenerateToken(Id uint, email, FullName string) (TokenInfo, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(24 * time.Hour)

	claims := CustomClaims{
		Id:       Id,
		Email:    email,
		FullName: FullName,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "Togo",
		},
	}

	token, err := j.CreateToken(claims)
	return TokenInfo{
		Token:     token,
		ExpiredAt: expireTime,
	}, err
}

func (j *JWT) GenerateTokenShort(Id uint) (TokenInfo, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(24 * time.Hour)

	claims := CustomClaimsShort{
		Id: Id,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "Togo",
		},
	}

	token, err := j.CreateTokenShort(claims)
	return TokenInfo{
		Token:     token,
		ExpiredAt: expireTime,
	}, err
}

func (j *JWT) CreateToken(claims CustomClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(j.SigningKey)
}

func (j *JWT) CreateTokenShort(claims CustomClaimsShort) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(j.SigningKey)
}

func (j *JWT) ParseToken(tokenString string) (*CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return j.SigningKey, nil
	})
	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, TokenMalformed
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				return nil, TokenExpired
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, TokenNotValidYet
			} else {
				return nil, TokenInvalid
			}
		}
	}

	if token != nil {
		if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
			return claims, nil
		}
		return nil, TokenInvalid
	} else {
		return nil, TokenInvalid
	}
}

func GetMsg(code int) string {
	msg, _ := MsgFlags[code]
	return msg
}

func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("Authorization")
		if token == "" {
			Response(c, http.StatusUnauthorized, false, GetMsg(ERROR_AUTH_TOKEN), gin.H{
				"reload": true,
			}, nil)
			c.Abort()
			return
		}

		j := NewJWT()
		claims, err := j.ParseToken(token)
		if err != nil {
			if err == TokenExpired {
				Response(c, http.StatusUnauthorized, false, GetMsg(ERROR_AUTH_CHECK_TOKEN_TIMEOUT), gin.H{
					"reload": true,
				}, nil)

				c.Abort()
				return
			}
			Response(c, http.StatusUnauthorized, false, GetMsg(ERROR_AUTH_CHECK_TOKEN_FAIL), gin.H{
				"reload": true,
			}, nil)
			c.Abort()
			return
		}
		c.Set("claims", claims)
		c.Next()
	}
}

func Response(c *gin.Context, httpCode int, success bool, message string, data interface{}, err error) {
	c.JSON(httpCode, gin.H{
		"success": success,
		"message": message,
		"data":    data,
		"error":   err,
	})
	return
}
