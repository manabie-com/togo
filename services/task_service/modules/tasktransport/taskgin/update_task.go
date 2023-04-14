package taskgin

import (
	"net/http"
	"strconv"

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

func UpdateTask(sc goservice.ServiceContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		user := c.MustGet("current_user").(*sdkcm.SimpleUser)

		var data taskmodel.TaskUpdate

		id, err := strconv.Atoi(c.Param("id"))

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})

			return
		}

		if err = c.ShouldBind(&data); err != nil {
			panic(err)
		}

		db := sc.MustGet(common.DBMain).(*gorm.DB)
		store := taskstorage.NewSQLStore(db)
		repo := taskrepo.NewRepo(store)

		hdl := taskhandler.NewUpdateTaskHdl(repo, user)
		if err = hdl.Response(c.Request.Context(), id, &data); err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, sdkcm.SimpleSuccessResponse("ok"))
	}
}
