package taskgin

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	goservice "github.com/phathdt/libs/go-sdk"
	"github.com/phathdt/libs/go-sdk/sdkcm"
	"github.com/phathdt/libs/togo_appgrpc"
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

		userService := sc.MustGet(common.PluginGrpcUserClient).(togo_appgrpc.UserClient)

		db := sc.MustGet(common.DBMain).(*gorm.DB)
		redisClient := sc.MustGet(common.PluginRedis).(*redis.Client)
		store := taskstorage.NewSQLStore(db)
		redisStore := taskstorage.NewRedisStore(redisClient)
		repo := taskrepo.NewRepo(store, redisStore)
		hdl := taskhandler.NewCreateTaskHdl(repo, userService, user)

		if err := hdl.Response(c.Request.Context(), &data); err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, sdkcm.SimpleSuccessResponse("ok"))
	}
}
