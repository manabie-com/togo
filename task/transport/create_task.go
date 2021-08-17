package transport

import (
	"github.com/gin-gonic/gin"
	appctx "github.com/manabie-com/togo/app_ctx"
	"github.com/manabie-com/togo/shared"
	"github.com/manabie-com/togo/task/dto"
	service2 "github.com/manabie-com/togo/task/service"
	"github.com/manabie-com/togo/task/storage"
	storage2 "github.com/manabie-com/togo/user/storage"
	"net/http"
)

func CreateTask(appCtx appctx.IAppCtx) gin.HandlerFunc {
	return func(c *gin.Context) {
		var data dto.CreateTaskRequest

		if err := c.ShouldBind(&data); err != nil {
			panic(shared.ErrInvalidRequest(err))
		}

		if err := data.Validate(); err != nil {
			panic(shared.ErrInvalidRequest(err))
		}

		requester := c.MustGet("current_user").(shared.Requester)
		db := appCtx.GetDbConn()
		tStore := storage.NewTaskStorage(db)
		uStore := storage2.NewUserStorage(db)
		service := service2.NewCreateTaskService(uStore, tStore, requester)

		if err := service.CreateTask(c.Request.Context(), &data); err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, shared.SimpleResponse(true))
		return
	}
}
