package ginuser

import (
	"github.com/japananh/togo/common"
	"github.com/japananh/togo/component"
	"github.com/japananh/togo/component/hasher"
	"github.com/japananh/togo/modules/user/userbiz"
	"github.com/japananh/togo/modules/user/usermodel"
	"github.com/japananh/togo/modules/user/userstorage"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Register(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var data usermodel.UserCreate

		if err := c.ShouldBind(&data); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		db := appCtx.GetMainDBConnection()
		store := userstorage.NewSQLStore(db)
		md5 := hasher.NewMd5Hash()
		biz := userbiz.NewRegisterBiz(store, md5)

		if err := biz.Register(c.Request.Context(), &data); err != nil {
			panic(err)
		}

		data.Mask()

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(data.FakeId.String()))
	}
}
