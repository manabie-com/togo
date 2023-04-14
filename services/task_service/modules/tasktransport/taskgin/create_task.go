package taskgin

import (
	"net/http"

	"github.com/gin-gonic/gin"
	goservice "github.com/phathdt/libs/go-sdk"
	"github.com/phathdt/libs/go-sdk/sdkcm"
	"gorm.io/gorm"
	"task_service/common"
	"task_service/modules/taskhandler"
	"task_service/modules/taskmodel"
	"task_service/modules/taskrepo"
	"task_service/modules/taskstorage"
)

func CreateTask(sc goservice.ServiceContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		user := c.MustGet("current_user").(*sdkcm.SimpleUser)

		var data taskmodel.TaskCreate

		if err := c.ShouldBind(&data); err != nil {
			panic(err)
		}

		data.UserId = user.ID

		db := sc.MustGet(common.DBMain).(*gorm.DB)
		store := taskstorage.NewSQLStore(db)
		repo := taskrepo.NewRepo(store)
		hdl := taskhandler.NewCreateTaskHdl(repo)

		if err := hdl.Response(c.Request.Context(), &data); err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, sdkcm.SimpleSuccessResponse("ok"))
	}
}
