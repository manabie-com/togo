package gintask

import (
	"github.com/gin-gonic/gin"
	"github.com/japananh/togo/common"
	"github.com/japananh/togo/component"
	"github.com/japananh/togo/modules/task/taskbiz"
	"github.com/japananh/togo/modules/task/taskmodel"
	"github.com/japananh/togo/modules/task/taskrepo"
	"github.com/japananh/togo/modules/task/taskstorage"
	"github.com/japananh/togo/modules/user/userstorage"
	"net/http"
)

func CreateTask(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var data taskmodel.TaskCreate

		if err := c.ShouldBind(&data); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		// get created_by from current user id
		requester := c.MustGet(common.CurrentUser).(common.Requester)
		data.CreatedBy = requester.GetUserId()

		// validate client input
		if err := data.Validate(); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		db := appCtx.GetMainDBConnection()
		store := taskstorage.NewSQLStore(db)
		userStore := userstorage.NewSQLStore(db)
		repo := taskrepo.NewCreateTaskRepo(store, userStore)
		biz := taskbiz.NewCreateTaskBiz(repo)

		if err := biz.CreateTask(c.Request.Context(), &data); err != nil {
			panic(err)
		}

		data.Mask()

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(data.FakeId.String()))
	}
}
