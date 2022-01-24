package rest

import (
	"context"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"togo/common"
	"togo/module/task/handler"
	"togo/module/task/model"
	"togo/module/task/repo"
	"togo/module/task/store/mysql"
	mysql2 "togo/module/userconfig/store/mysql"
	mysql3 "togo/module/usertask/store/mysql"

	"togo/server"
)

func CreateTasks(sc server.ServerContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := context.Background()
		var params model.CreateTasksParams
		if err := c.ShouldBind(&params); err != nil {
			c.AbortWithStatusJSON(common.BAD_REQUEST_STATUS, common.ResponseError(common.BAD_REQUEST_STATUS, err.Error()))
			return
		}

		if params.UserId == nil {
			c.AbortWithStatusJSON(common.BAD_REQUEST_STATUS, common.ResponseError(common.BAD_REQUEST_STATUS, "UserIsNull"))
			return
		}

		db := sc.GetService("MYSQL").(*gorm.DB)

		taskStore := mysql.NewTaskSQL(db)
		userStore := mysql2.NewUserConfigSQL(db)
		userTaskStore := mysql3.NewUserTaskSQL(db)
		taskRepo := repo.NewCreateTaskRepo(userStore, taskStore, userTaskStore)
		taskHdl := handler.NewCreateTaskHdl(taskRepo)

		if err := taskHdl.CreateTasks(ctx, *params.UserId, params.Tasks); err != nil {
			c.AbortWithStatusJSON(common.BAD_REQUEST_STATUS, common.ResponseError(common.BAD_REQUEST_STATUS, err.Error()))
			return
		}

		c.JSON(common.SUCCESS_STATUS, common.ResponseSuccess(common.SUCCESS_STATUS, nil))
	}
}