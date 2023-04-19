package services

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
	"time"
	"togo/pkg/models"
	"togo/pkg/utils"
)

func (s *Server) GetTaskByDate(ctx *gin.Context) {
	createdDate := ctx.Query("createdDate")
	userId, _ := ctx.Get("userId")

	var task []models.Task
	if result := s.H.DB.Where(&models.Task{
		UserId:      userId.(int64),
		CreatedDate: createdDate,
	}).Find(&task); result.Error != nil {
		ctx.JSON(http.StatusBadRequest, result.Error.Error())
		return
	}

	var response models.Response[models.Task]
	response.Status = http.StatusOK
	response.Message = "OK"
	response.Data = task
	ctx.JSON(http.StatusOK, response)
}

func (s *Server) CreateTask(ctx *gin.Context) {
	request := models.TaskRequest{}
	if err := ctx.BindJSON(&request); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	userIdJwt, _ := ctx.Get("userId")
	userId := userIdJwt.(int64)
	keyRedisCache := fmt.Sprintf("USER_LIMIT_TASK_%v", userIdJwt)
	result, _ := s.Redis.Exists(context.Background(), keyRedisCache).Result()
	if result == 1 {
		ctx.JSON(http.StatusBadRequest, "Limit create task by date.")
		return
	}

	var countTask models.CountTask
	var createDateNow = time.Now().Format("2006-01-02")
	query := s.H.DB.Table("users").Select("users.limit_task , count(*) as number_task")
	query = query.Joins(" left join tasks on users.id = tasks.user_id ")
	query = query.Where(" users.id = ? and tasks.created_date= ?", userId, createDateNow).Group("users.limit_task")
	query = query.Scan(&countTask)

	if countTask.LimitTask > 0 && countTask.NumberTask >= countTask.LimitTask {
		ctx.JSON(http.StatusBadRequest, "Limit create task by date.")
		go func() {
			nowDate := time.Now()
			endDate := utils.EndDate(nowDate)
			secondExpire := (endDate.UTC().UnixMilli() - nowDate.UTC().UnixMilli()) / 1000
			err := s.Redis.Set(context.Background(), keyRedisCache, "true", time.Duration(secondExpire)*time.Second).Err()
			if err != nil {
				fmt.Println("Set cache redis limit user task fail")
			}
		}()
		return
	}

	var task models.Task
	task.UserId = userId
	task.TaskId = uuid.New().String()
	task.CreatedDate = createDateNow
	task.EventTime = time.Now().Format("2006-01-02 15:04:05")
	task.Content = request.Content

	s.H.DB.Create(&task)

	var response models.Response[models.Task]
	response.Status = http.StatusOK
	response.Message = "Create task success"
	response.Data = []models.Task{task}
	ctx.JSON(http.StatusOK, response)
}
