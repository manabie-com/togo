package request

import (
	"github.com/gin-gonic/gin"
	"github.com/shanenoi/togo/config"
	"strings"
)

func GetHeader(ctx *gin.Context, key string) string {
	return ctx.Request.Header.Get(key)
}

func GetAuthToken(ctx *gin.Context) (token string)  {
	authString := GetHeader(ctx, config.HEADER_AUTH)
	if len(authString) > 0 {
		values := strings.Split(authString, " ")
		lenValues := len(values)
		if lenValues > 0 {
			token = values[lenValues-1]
		}
	}
	return
}
