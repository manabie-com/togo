package http

import (
	"ansidev.xyz/pkg/log"
	"github.com/ansidev/togo/auth/dto"
	"github.com/ansidev/togo/auth/service"
	"github.com/ansidev/togo/errs"
	"github.com/gin-gonic/gin"
	"net/http"
)

func registerRoutes(router *gin.Engine, authController *authController) {
	v1 := router.Group("/auth/v1")

	v1.POST("/login", authController.Login)
}

func NewAuthController(router *gin.Engine, authService service.IAuthService) {
	controller := &authController{authService}
	registerRoutes(router, controller)
}

type authController struct {
	authService service.IAuthService
}

func (ctrl *authController) Login(ctx *gin.Context) {
	var credential dto.UsernamePasswordCredential

	if err := ctx.ShouldBindJSON(&credential); err != nil {
		ctx.Error(err)
		return
	}

	token, err := ctrl.authService.Login(credential)

	if err != nil {
		log.Debug(err)
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, errs.ErrorResponse{
			Code:    errs.ErrCodeUnauthorized,
			Message: http.StatusText(http.StatusUnauthorized),
			Error:   http.StatusText(http.StatusUnauthorized),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"token": token})
}
