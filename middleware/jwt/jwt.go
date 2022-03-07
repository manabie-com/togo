package jwt

import (
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"

	"github.com/khoale193/togo/pkg/e"
	"github.com/khoale193/togo/pkg/util"
)

// JWT is jwt middleware
func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		var data interface{}
		code := e.SUCCESS
		token := c.Request.Header.Get(e.UserAuth)
		if token == "" {
			code = e.INVALID_PARAMS
		} else {
			print(strings.ToUpper(token[0:6]))
			if len(token) > 6 && strings.ToUpper(token[0:7]) == "BEARER " {
				token = token[7:]
			}
			_, err := util.ParseToken(token)
			if err != nil {
				switch err.(*jwt.ValidationError).Errors {
				case jwt.ValidationErrorExpired:
					code = e.ERROR_AUTH_CHECK_TOKEN_TIMEOUT
				case jwt.ValidationErrorSignatureInvalid:
					code = e.ERROR_UNAUTHORIZES
				default:
					code = e.ERROR_AUTH_CHECK_TOKEN_FAIL
				}
			}
		}
		if code != e.SUCCESS {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": code,
				"msg":  e.GetMsg(code),
				"data": data,
			})
			c.Abort()
			return
		}
		c.Next()
	}
}
