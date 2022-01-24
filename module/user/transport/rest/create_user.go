package rest

import (
	"context"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"togo/common"
	"togo/module/user/handler"
	"togo/module/user/model"
	"togo/module/user/repo"
	"togo/module/user/store"
	"togo/module/userconfig/store/mysql"

	"togo/server"
)

func CreateUser(sc server.ServerContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := context.Background()
		var data model.CreateUser
		if err := c.ShouldBind(&data); err != nil {
			c.AbortWithStatusJSON(common.BAD_REQUEST_STATUS, common.ResponseError(common.BAD_REQUEST_STATUS, err.Error()))
		}

		db := sc.GetService("MYSQL").(*gorm.DB)

		userStore := store.NewUserSQL(db)
		userCfgStore := mysql.NewUserConfigSQL(db)
		userRepo := repo.NewCreateUserRepo(userStore, userCfgStore)
		userHdl := handler.NewCreateUserHdl(userRepo)

		if err := userHdl.CreateUser(ctx, &data); err != nil {
			c.AbortWithStatusJSON(common.BAD_REQUEST_STATUS, "")
		}

		c.JSON(common.SUCCESS_STATUS, common.ResponseSuccess(common.SUCCESS_STATUS, data))
	}
}