package transport

import (
	"github.com/gin-gonic/gin"
	appctx "github.com/manabie-com/togo/app_ctx"
	"github.com/manabie-com/togo/auth/dto"
	service2 "github.com/manabie-com/togo/auth/service"
	"github.com/manabie-com/togo/auth/storage"
	"github.com/manabie-com/togo/shared"
	"net/http"
)

func UserLogin(appctx appctx.IAppCtx) gin.HandlerFunc {
	return func(c *gin.Context) {
		var data dto.LoginRequest

		if err := c.ShouldBind(&data); err != nil {
			panic(shared.ErrInvalidRequest(err))
		}

		if err := data.Validate(); err != nil {
			panic(err)
		}

		db := appctx.GetDbConn()

		store := storage.NewAuthStorage(db)
		service := service2.NewAuthService(store, appctx.GetTokenProvider())
		loginResponse, err := service.Login(c.Request.Context(), &data)
		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, shared.SimpleResponse(loginResponse))
		return
	}
}
