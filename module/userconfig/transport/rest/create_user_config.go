package rest

import (
	"context"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"togo/common"
	"togo/module/userconfig/handler"
	"togo/module/userconfig/model"
	"togo/module/userconfig/repo"
	"togo/module/userconfig/store/mysql"

	"togo/server"
)

func CreateUserConfig(sc server.ServerContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := context.Background()
		var data model.CreateUserConfig
		if err := c.ShouldBind(&data); err != nil {
			c.AbortWithStatusJSON(common.BAD_REQUEST_STATUS, common.ResponseError(common.BAD_REQUEST_STATUS, err.Error()))
		}

		db := sc.GetService("MYSQL").(*gorm.DB)

		userStore := mysql.NewUserConfigSQL(db)
		userRepo := repo.NewCreateUserConfigRepo(userStore)
		userHdl := handler.NewCreateUserConfigHdl(userRepo)

		if err := userHdl.CreateUserConfig(ctx, &data); err != nil {
			c.AbortWithStatusJSON(common.BAD_REQUEST_STATUS, "")
		}

		c.JSON(common.SUCCESS_STATUS, common.ResponseSuccess(common.SUCCESS_STATUS, data))
	}
}