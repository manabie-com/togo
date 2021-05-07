package user

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"manabie-com/togo/entity"
	pkg_logrus "manabie-com/togo/pkg/logger"
	pkg_rd "manabie-com/togo/pkg/rd"
	"manabie-com/togo/util"
	"net/http"
	"strconv"
	"time"
)

func mergeRedisKey(userID, createdDate string) string {
	return userID + "_" + createdDate
}

func AddTask(c *gin.Context) {
	var f entity.Task
	if err := c.ShouldBindJSON(&f); err != nil {
		util.AbortJSONBadRequest(c)
		return
	}
	userID := util.TokenUserID(c)

	now := time.Now()
	createdDate := now.Format(util.DefaultTimeFormat)
	todayRedisKey := mergeRedisKey(userID, createdDate)
	var isReached bool

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// First, we can reduce query to DB by check the daily limit is reached by user_id with redis
	result, err := pkg_rd.RdConn().Get(ctx, todayRedisKey).Result()
	if err == redis.Nil || err == nil {
		isReached, _ = strconv.ParseBool(result)
	} else {
		util.AbortUnexpected(c, util.ERR_CODE_DB_ISSUE, "can't check max to do")
		return
	}

	if isReached {
		util.AbortUnauthorized(c, util.ERR_CODE_USER_REACH_LIMIT_TASK, "The daily limit is reached")
		return
	}

	task := entity.Task{
		ID:          uuid.New().String(),
		Content:     f.Content,
		UserID:      userID,
		CreatedDate: createdDate,
	}

	maxTodo := util.TokenMaxTodo(c)
	var count int64
	err = entity.Db().Transaction(func(tx *gorm.DB) error {

		if err := tx.Model(entity.Task{}).Where("user_id = ? AND created_date = ?", userID, createdDate).Count(&count).Error; err != nil {
			pkg_logrus.Lgrus.Error("can't count task: ", userID, createdDate)
			return err
		}

		if count < int64(maxTodo) {
			if err := tx.Model(entity.Task{}).Create(&task).Error; err != nil {
				pkg_logrus.Lgrus.Error("create task:", err.Error())
				return err
			}
		} else {
			isReached = true
		}
		return nil
	})

	if err != nil {
		util.AbortUnexpected(c, util.ERR_CODE_DB_ISSUE, err.Error())
		return
	}

	if isReached {
		util.AbortUnauthorized(c, util.ERR_CODE_USER_REACH_LIMIT_TASK, "The daily limit is reached")
		return
	}

	if count == int64(maxTodo)-1 {
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()
		pkg_rd.RdConn().Set(ctx, todayRedisKey, true, time.Hour*24)
	}

	c.JSON(http.StatusOK, gin.H{"detail": "success"})
}
