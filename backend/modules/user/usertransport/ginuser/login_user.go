package ginuser

import (
	"net/http"

	"github.com/gin-gonic/gin"
	goservice "github.com/phathdt/libs/go-sdk"
	"github.com/phathdt/libs/go-sdk/plugin/tokenprovider"
	"github.com/phathdt/libs/go-sdk/sdkcm"
	"gorm.io/gorm"
	"togo/common"
	"togo/modules/user/userhandler"
	"togo/modules/user/usermodel"
	"togo/modules/user/userrepo"
	"togo/modules/user/userstorage"
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
		repo := userrepo.NewRepo(store)
		md5 := common.NewMd5Hash()

		hdl := userhandler.NewLoginHdl(repo, tokenProvider, md5, 60*60*24*30)
		account, err := hdl.Response(c.Request.Context(), &loginUserData)
		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, sdkcm.SimpleSuccessResponse(account))
	}
}
