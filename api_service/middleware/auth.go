package middleware

import (
	"api_service/connection"
	"api_service/proto"
	"net/http"

	"github.com/gin-gonic/gin"
)

//TokenAuth ...
func TokenAuth(ctx *gin.Context) {
	err := CheckTokenValidation(ctx)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, "Login Required!")
		ctx.Abort()
	}
	ctx.Next()
}

//CheckTokenValidation ...
func CheckTokenValidation(ctx *gin.Context) error {

	cookie, err := GetTokenCookie(ctx)
	if err != nil {
		return err
	}

	tokenString := cookie.Value
	sc := connection.DialToSessionServiceServer()
	_, checkErr := sc.ClientSessionService.CheckToken(ctx, &proto.TokenString{
		Token: tokenString,
	})

	if checkErr != nil {
		return checkErr
	}

	return nil
}

//GetTokenCookie ...
func GetTokenCookie(ctx *gin.Context) (*http.Cookie, error) {

	cookie, err := ctx.Request.Cookie("token")
	if err != nil {
		if err == http.ErrNoCookie {
			return nil, err
		}
		return nil, err
	}

	return cookie, nil
}
