package ginuser

import (
	"net/http"

	"github.com/gin-gonic/gin"
	goservice "github.com/phathdt/libs/go-sdk"
	"github.com/phathdt/libs/go-sdk/sdkcm"
	"gorm.io/gorm"
	"user_service/common"
	"user_service/modules/userhandler"
	"user_service/modules/usermodel"
	"user_service/modules/userstorage"
)

func Register(sc goservice.ServiceContext) func(ctx *gin.Context) {
	return func(c *gin.Context) {

		var data usermodel.UserCreate

		if err := c.ShouldBind(&data); err != nil {
			panic(err)
		}

		db := sc.MustGet(common.DBMain).(*gorm.DB)

		store := userstorage.NewSQLStore(db)
		md5 := common.NewMd5Hash()

		hdl := userhandler.NewRegisterUserHdl(store, md5)

		if err := hdl.Response(c.Request.Context(), &data); err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, sdkcm.SimpleSuccessResponse("ok"))
	}
}
