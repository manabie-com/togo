package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/manabie/project/pkg/jwt"
	"strings"
)

type accessController struct {
	jwt jwt.TokenUser
}

type AccessController interface {
	Authenticate() gin.HandlerFunc
}

func NewAccessController(jwt jwt.TokenUser) AccessController {
	return &accessController{
		jwt:jwt,
	}
}

func(a *accessController) Authenticate() gin.HandlerFunc {
	return func(c *gin.Context) {
		bearToken:=c.Request.Header.Get("Authorization")
		if bearToken == ""{
			c.AbortWithStatusJSON(401, gin.H{
				"message": "Authentication failure: Token not provided",
			})
			return
		}
		strArr := strings.Split(bearToken, " ")
		message,err:=a.jwt.ParseToken(strArr[1])
		if err!=nil{
			c.AbortWithStatusJSON(400, gin.H{
				"message": message,
			})
			return
		}
		c.Next()
	}
}
