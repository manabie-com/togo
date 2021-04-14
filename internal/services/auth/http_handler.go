package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Handler(
	router *gin.Engine,
	authSrv Service,
) {
	router.POST("/login", loginHandler(authSrv))
}

func loginHandler(authSrv Service) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		userName := ctx.PostForm("username")
		password := ctx.PostForm("password")
		token, err := authSrv.login(ctx, userName, password)
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"status": "ERROR",
				"error": map[string]string{
					"code":    "INVALID_REQUEST",
					"message": err.Error(),
				},
			})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"status": "OK",
			"token":  token,
		})
	}
}
