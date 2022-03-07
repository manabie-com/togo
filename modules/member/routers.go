package member

import (
	"github.com/gin-gonic/gin"

	"github.com/khoale193/togo/modules/member/handler"
)

func MemberRouter(router *gin.RouterGroup) {
	router.POST("/sign-in", handler.SignIn)
}
