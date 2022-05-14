package ginuser

import (
	"github.com/japananh/togo/common"
	"github.com/japananh/togo/component"
	"github.com/japananh/togo/component/hasher"
	"github.com/japananh/togo/component/tokenprovider/jwt"
	"github.com/japananh/togo/modules/user/userbiz"
	"github.com/japananh/togo/modules/user/usermodel"
	"github.com/japananh/togo/modules/user/userstorage"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Login(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var data usermodel.UserLogin

		if err := c.ShouldBind(&data); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		db := appCtx.GetMainDBConnection()
		store := userstorage.NewSQLStore(db)
		tokenProvider := jwt.NewTokenJWTProvider(appCtx.SecretKey())
		md5 := hasher.NewMd5Hash()
		tokenConfig := appCtx.GetTokenConfig()

		biz := userbiz.NewLoginBiz(store, tokenProvider, md5, tokenConfig)

		account, err := biz.Login(c.Request.Context(), &data)
		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(account))
	}
}
