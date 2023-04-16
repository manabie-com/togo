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
	"togo/modules/user/userrepo"
	"togo/modules/user/userstorage"
)

func CreateTask(sc goservice.ServiceContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		user := c.MustGet("current_user").(*sdkcm.SimpleUser)

		var data taskmodel.TaskCreate

		if err := c.ShouldBind(&data); err != nil {
			panic(err)
		}

		db := sc.MustGet(common.DBMain).(*gorm.DB)

		redisClient := sc.MustGet(common.PluginRedis).(*redis.Client)
		store := taskstorage.NewSQLStore(db)
		redisStore := taskstorage.NewRedisStore(redisClient)
		repo := taskrepo.NewRepo(store, redisStore)

		userStore := userstorage.NewSQLStore(db)
		userRepo := userrepo.NewRepo(userStore)
		hdl := taskhandler.NewCreateTaskHdl(repo, userRepo, user)

		if err := hdl.Response(c.Request.Context(), &data); err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, sdkcm.SimpleSuccessResponse("ok"))
	}
}
