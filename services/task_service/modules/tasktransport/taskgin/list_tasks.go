package taskgin

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
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

		fmt.Printf("query = %#v\n", queryString.Filter)

		db := sc.MustGet(common.DBMain).(*gorm.DB)
		redisClient := sc.MustGet(common.PluginRedis).(*redis.Client)
		store := taskstorage.NewSQLStore(db)
		redisStore := taskstorage.NewRedisStore(redisClient)
		repo := taskrepo.NewRepo(store, redisStore)
		hdl := taskhandler.NewListTaskHdl(repo)

		tasks, err := hdl.Response(c.Request.Context(), &queryString.Filter, &queryString.Paging)
		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, sdkcm.ResponseWithPaging(tasks, queryString.Paging, queryString.Filter))
	}
}
