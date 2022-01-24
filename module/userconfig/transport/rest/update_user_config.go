package rest

import (
	"context"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"strconv"
	"togo/common"
	"togo/module/userconfig/handler"
	"togo/module/userconfig/model"
	"togo/module/userconfig/repo"
	"togo/module/userconfig/store/mysql"

	"togo/server"
)

func UpdateUserConfig(sc server.ServerContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := context.Background()
		var data model.UpdateUserConfig
		id := c.Param("id")
		if id == "" {
			c.AbortWithStatusJSON(common.BAD_REQUEST_STATUS, common.ResponseError(common.BAD_REQUEST_STATUS, "UserIsNull"))
			return
		}

		userId, err := strconv.ParseInt(id, 10, 32)
		if err != nil {
			c.AbortWithStatusJSON(common.BAD_REQUEST_STATUS, common.ResponseError(common.BAD_REQUEST_STATUS, err.Error()))
			return
		}

		if err := c.ShouldBind(&data); err != nil {
			c.AbortWithStatusJSON(common.BAD_REQUEST_STATUS, common.ResponseError(common.BAD_REQUEST_STATUS, err.Error()))
			return
		}

		db := sc.GetService("MYSQL").(*gorm.DB)

		userStore := mysql.NewUserConfigSQL(db)
		userRepo := repo.NewUpdateUserConfigRepo(userStore)
		userHdl := handler.NewUpdateUserConfigHdl(userRepo)

		if err := userHdl.UpdateUserConfig(ctx, uint(userId), &data); err != nil {
			c.AbortWithStatusJSON(common.BAD_REQUEST_STATUS, "")
			return
		}

		c.JSON(common.SUCCESS_STATUS, common.ResponseSuccess(common.SUCCESS_STATUS, data))
	}
}