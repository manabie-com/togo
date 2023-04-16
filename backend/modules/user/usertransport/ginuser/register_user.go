package ginuser

import (
	"net/http"

	"github.com/gin-gonic/gin"
	goservice "github.com/phathdt/libs/go-sdk"
	"github.com/phathdt/libs/go-sdk/sdkcm"
	"gorm.io/gorm"
	"togo/common"
	"togo/modules/user/userhandler"
	"togo/modules/user/usermodel"
	"togo/modules/user/userrepo"
	"togo/modules/user/userstorage"
)

func Register(sc goservice.ServiceContext) gin.HandlerFunc {
	return func(c *gin.Context) {

		var data usermodel.UserCreate

		if err := c.ShouldBind(&data); err != nil {
			panic(err)
		}

		db := sc.MustGet(common.DBMain).(*gorm.DB)

		store := userstorage.NewSQLStore(db)
		repo := userrepo.NewRepo(store)
		md5 := common.NewMd5Hash()

		hdl := userhandler.NewRegisterUserHdl(repo, md5)

		if err := hdl.Response(c.Request.Context(), &data); err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, sdkcm.SimpleSuccessResponse("ok"))
	}
}
