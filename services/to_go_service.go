package services

import (
	"errors"
	"github.com/namnhatdoan/togo/constants"
	"github.com/namnhatdoan/togo/db"
	"github.com/namnhatdoan/togo/models"
	"github.com/namnhatdoan/togo/settings"
	"github.com/namnhatdoan/togo/utils"
	"gorm.io/gorm"
	"time"
)

type ToGoServiceImpl struct {}

func (s *ToGoServiceImpl) AddNewTask(email, task string) (*models.Tasks, error) {
	config, err := s.GetOrCreateConfigForUpdate(email, utils.GetCurrentDate())
	if err != nil {
		return nil, err
	}
	if config.Current >= config.Limit {
		return nil, errors.New(constants.ExceedTaskPerDayLimit)
	}
	defer db.GetDB().Updates(config)

	newTask := &models.Tasks{
		Email: email,
		Task: task,
	}
	res := db.GetDB().Create(newTask)
	if res.Error != nil {
		settings.GetLogger().WithError(res.Error).Error("Create new task fail")
		return nil, errors.New(constants.CreateNewTaskFail)
	}
	config.Current += 1
	return newTask, nil
}

func (s *ToGoServiceImpl) GetOrCreateConfigForUpdate(email string, date time.Time) (*models.UserConfigs, error) {
	config, err := db.GetConfigForUpdate(email, date)

	if err == nil {
		// Config found
		return config, nil
	} else if err == gorm.ErrRecordNotFound {
		// Create config if not found
		log.WithField("email", email).WithField("date", date).Info("Not found config")
		if _, err := db.UpsertUserConfig(email, constants.DefaultLimitTaskPerDay, utils.GetCurrentDate()); err != nil {
			return nil, err
		}
	} else {
		return nil, err
	}

	return db.GetConfigForUpdate(email, date)
}

func (s *ToGoServiceImpl) SetUserConfig(email string, limit int8, date time.Time) (*models.UserConfigs, error) {
	return db.UpsertUserConfig(email, limit, date)
}
