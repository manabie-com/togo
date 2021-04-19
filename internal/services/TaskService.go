package services

import (
	"time"
	"net/http"
	"errors"
	// "fmt"
	"log"
	"github.com/manabie-com/togo/internal/storages"
	"github.com/manabie-com/togo/internal/utils"
	
)

type ITaskService interface {
	StoreFromService(task storages.Task) (error, int)
	FindByIdAndTimeFromService(user_id, time string) ([]storages.Task, error)
}

type TaskService struct {
	TaskRepo storages.ITaskRepo
	QuotaRepo storages.IQuotaRepo
}

func (service *TaskService) StoreFromService(task storages.Task) (error, int) {

	q, _ := service.QuotaRepo.FindByUserIdFromRepo(task.UserID)

	if q.UserID == "" {
		log.Println("InitQuota")
		q = service.QuotaRepo.InitQuota(task.UserID)
	}

	lessThanLimit := utils.LessThanLimit(q.StartTime, 24*time.Hour)

	if lessThanLimit && q.Quota <= 0 {
		return errors.New("Too Many Requests"), http.StatusTooManyRequests
	}

	if !lessThanLimit {
		q.StartTime = utils.GetStartTime()
		q.Quota = storages.QUOTA_LIMIT
	}
	q.Quota = q.Quota-1	

	err := service.TaskRepo.StoreFromRepo(task)
	service.QuotaRepo.ReplaceFromRepo(q)

	return err, http.StatusInternalServerError
}

func (service *TaskService) FindByIdAndTimeFromService(user_id, time string) ([]storages.Task, error) {
	tasks, err := service.TaskRepo.FindByIdAndTimeFromRepo(user_id, time)
	if err != nil || len(tasks) == 0 {
		return []storages.Task{}, err
	}

	return tasks, err
}

