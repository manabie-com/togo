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

func ListTasks(sc goservice.ServiceContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		user := c.MustGet("current_user").(*sdkcm.SimpleUser)

		var queryString struct {
			sdkcm.Paging
			taskmodel.Filter
		}

		if err := c.ShouldBind(&queryString); err != nil {
			panic(err)
		}

		queryString.Paging.FullFill()

		queryString.Filter.UserId = user.ID

		db := sc.MustGet(common.DBMain).(*gorm.DB)
		store := taskstorage.NewSQLStore(db)
		repo := taskrepo.NewRepo(store)
		hdl := taskhandler.NewListTaskHdl(repo)

		tasks, err := hdl.Response(c.Request.Context(), &queryString.Filter, &queryString.Paging)
		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, sdkcm.ResponseWithPaging(tasks, queryString.Paging, queryString.Filter))
	}
}
