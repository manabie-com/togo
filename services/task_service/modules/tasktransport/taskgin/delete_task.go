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
	"task_service/modules/taskrepo"
	"task_service/modules/taskstorage"
)

func DeleteTask(sc goservice.ServiceContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		user := c.MustGet("current_user").(*sdkcm.SimpleUser)

		id, err := strconv.Atoi(c.Param("id"))

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})

			return
		}

		db := sc.MustGet(common.DBMain).(*gorm.DB)
		store := taskstorage.NewSQLStore(db)
		repo := taskrepo.NewRepo(store)

		hdl := taskhandler.NewDeleteTaskHdl(repo, user)
		if err = hdl.Response(c.Request.Context(), id); err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, sdkcm.SimpleSuccessResponse("ok"))
	}
}
