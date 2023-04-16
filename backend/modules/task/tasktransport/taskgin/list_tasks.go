package taskgin

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	goservice "github.com/phathdt/libs/go-sdk"
	"github.com/phathdt/libs/go-sdk/sdkcm"
	"gorm.io/gorm"
	"togo/common"
	"togo/modules/task/taskhandler"
	"togo/modules/task/taskmodel"
	"togo/modules/task/taskrepo"
	"togo/modules/task/taskstorage"
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

		db := sc.MustGet(common.DBMain).(*gorm.DB)
		store := taskstorage.NewSQLStore(db)
		redisClient := sc.MustGet(common.PluginRedis).(*redis.Client)
		redisStore := taskstorage.NewRedisStore(redisClient)
		repo := taskrepo.NewRepo(store, redisStore)
		hdl := taskhandler.NewListTaskHdl(repo, user)

		tasks, err := hdl.Response(c.Request.Context(), &queryString.Filter, &queryString.Paging)
		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, sdkcm.ResponseWithPaging(tasks, queryString.Paging, queryString.Filter))
	}
}
