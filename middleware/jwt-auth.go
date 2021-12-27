package middleware

import (

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"log"
	"togo/db"
	"togo/form"
	"togo/model"
)

var IdentityKey = "username"

func PayloadFunc(data interface{}) jwt.MapClaims {
	if v, ok := data.(*model.User); ok {
		log.Println(v)
		return jwt.MapClaims{
			IdentityKey: v.Username,
		}
	}
	return jwt.MapClaims{}
}

func IdentityHandler(c *gin.Context) interface{} {
	claims := jwt.ExtractClaims(c)
	return &model.User{
		Username: claims[IdentityKey].(string),
	}
}

func Authenticator(c *gin.Context) (interface{}, error) {
	var loginVals form.Login
	if err := c.ShouldBind(&loginVals); err != nil {
		return "", jwt.ErrMissingLoginValues
	}
	username := loginVals.Username
	password := loginVals.Password

	userVals := model.User{}
	db.DB.First(&userVals, "username = ? AND password = ?", username, password)
	if username == userVals.Username && password == userVals.Password {
		return &model.User{
			Username: username,
		}, nil
	}

	return nil, jwt.ErrFailedAuthentication
}

func Authorizator(data interface{}, c *gin.Context) bool {
	claims := jwt.ExtractClaims(c)
	if v, ok := data.(*model.User); ok && v.Username == claims[IdentityKey] {
		return true
	}

	return false
}

func Unauthorized(c *gin.Context, code int, message string) {
	c.JSON(code, gin.H{
		"code":    code,
		"message": message,
	})
}

func GetUserFromCtx(c *gin.Context) model.User {
	claims := jwt.ExtractClaims(c)
	userVals := model.User{}
	db.DB.First(&userVals, "username = ?", claims[IdentityKey].(string))
	return userVals
}
