package libs

import (
	"os"
	"strings"

	"github.com/dgrijalva/jwt-go"

	"github.com/gin-gonic/gin"
)

// RecoverError func
func RecoverError(c *gin.Context) {
	if r := recover(); r != nil {
		responseData := gin.H{
			"status": 500,
			"msg":    r,
		}
		c.JSON(500, responseData)
		return
	}
}

// APIResponseData func
func APIResponseData(c *gin.Context, status int, responseData gin.H) {
	responseType := c.Request.Header.Get("ResponseType")
	if responseType == "application/xml" {
		c.XML(status, responseData)
	} else {
		c.JSON(status, responseData)
	}
}

func GetUserIDFromToken(token string) string {
	var userID string
	claims := make(jwt.MapClaims)
	t, err := jwt.ParseWithClaims(token, claims, func(*jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_TOKEN")), nil
	})
	if err == nil {
		if t.Valid {
			sUserID, ok := claims["user_id"].(string)
			if ok {
				userID = sUserID
			}
		}
	}
	return userID
}

// GetQueryParam func
func GetQueryParam(param string, c *gin.Context) (string, bool) {
	vParam, sParam := c.GetQuery(param)
	vParamLower, sParamLower := c.GetQuery(strings.ToLower(param))
	if sParam {
		return vParam, sParam
	}
	if sParamLower {
		return vParamLower, sParamLower
	}
	return "", false
}
