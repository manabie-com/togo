package tasklimiter

import (
	"fmt"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"manabie.com/togo"
)

type Config struct {
	Store togo.Store
	Cache togo.Cache
}

type taskLimiterServiceImpl struct {
	store togo.Store
	cache togo.Cache
}

func New(cfg Config) togo.TaskLimiterService {
	return &taskLimiterServiceImpl{store: cfg.Store, cache: cfg.Cache}
}

func (t *taskLimiterServiceImpl) CreateTask(rq togo.TaskCreationRequest) (togo.TaskCreationResponse, error) {
	entity, err := t.store.CreateTask(rq.TaskName, rq.Description)
	if err != nil {
		return togo.TaskCreationResponse{}, errors.Wrapf(err, "fail to create task")
	}

	return togo.TaskCreationResponse{Id: entity.Id, TaskName: entity.TaskName, Description: entity.Description, CreatedAt: entity.CreatedAt}, nil
}

func (t *taskLimiterServiceImpl) DoTask(userId int, taskId int) (string, error) {
	if numDo, err := t.cache.CheckIfCanDoTask(userId, taskId); err != nil {
		if errRollback := t.cache.RollbackIfCheckFail(userId, taskId); errRollback != nil {
			return "", errRollback
		}
		return "", err
	} else {
		return fmt.Sprintf("task %d has been done by user %d for the %d times", taskId, userId, numDo), nil
	}
}

func (t *taskLimiterServiceImpl) SetTaskLimit(userId int, taskId int, limit int) (string, error) {
	if err := t.store.SetTaskLimit(userId, taskId, limit); err != nil {
		return "", err
	}

	if err := t.cache.SetTaskLimit(userId, taskId, limit); err != nil {
		return "", err
	}

	return fmt.Sprintf("user %d now can do task %d - %d times per day", userId, taskId, limit), nil
}

func (t *taskLimiterServiceImpl) ResetDailyTask() (string, error) {
	taskLimitSettings, err := t.store.GetAllTaskLimitSetting()
	if err != nil {
		return "", errors.Wrapf(err, "fail to get all task limit setting")
	}
	log.Infof("get all task limit settings from database success %d", len(taskLimitSettings))

	if err = t.cache.ResetTaskForNewDay(taskLimitSettings); err != nil {
		return "", err
	}

	return "load all task limit setting to cache", nil
}
