package ginuser

import (
	"net/http"

	"github.com/gin-gonic/gin"
	goservice "github.com/phathdt/libs/go-sdk"
	"github.com/phathdt/libs/go-sdk/plugin/tokenprovider"
	"github.com/phathdt/libs/go-sdk/sdkcm"
	"gorm.io/gorm"
	"user_service/common"
	"user_service/modules/userhandler"
	"user_service/modules/usermodel"
	"user_service/modules/userstorage"
)

func Login(sc goservice.ServiceContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		db := sc.MustGet(common.DBMain).(*gorm.DB)
		tokenProvider := sc.MustGet(common.PluginJWT).(tokenprovider.Provider)

		var loginUserData usermodel.UserLogin

		if err := c.ShouldBind(&loginUserData); err != nil {
			panic(sdkcm.ErrInvalidRequest(err))
		}

		store := userstorage.NewSQLStore(db)
		md5 := common.NewMd5Hash()

		hdl := userhandler.NewLoginHdl(store, tokenProvider, md5, 60*60*24*30)
		account, err := hdl.Response(c.Request.Context(), &loginUserData)
		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, sdkcm.SimpleSuccessResponse(account))
	}
}
