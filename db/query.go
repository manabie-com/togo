package db

import (
	"errors"
	"github.com/namnhatdoan/togo/constants"
	"github.com/namnhatdoan/togo/models"
	"github.com/namnhatdoan/togo/settings"
	"github.com/namnhatdoan/togo/utils"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"time"
)

func AddNewTask(email, task string) (*models.Tasks, error) {
	config, err := GetConfig(email, utils.GetCurrentDate())
	if err != nil {
		return nil, err
	}
	if config.Current >= config.Limit {
		return nil, errors.New(constants.ExceedTaskPerDayLimit)
	}
	defer db.Updates(config)
	newTask := &models.Tasks{
		Email: email,
		Task: task,
	}
	res := db.Create(newTask)
	if res.Error != nil {
		settings.GetLogger().WithError(res.Error).Error("Create new task fail")
		return nil, res.Error
	}
	config.Current += 1
	return newTask, nil
}

func GetConfig(email string, date time.Time) (*models.UserConfigs, error) {
	config := &models.UserConfigs{}
	res := db.Clauses(clause.Locking{
		Strength: "UPDATE",
		Options: "NOWAIT",
	}).Where(&models.UserConfigs{
		Email: email,
		Date: date,
	}).First(config)

	if res.Error == gorm.ErrRecordNotFound{
		// Create config if not found
		log.WithField("email", email).WithField("date", date).Info("Not found config")
		if _, err := SetConfig(email, constants.DefaultLimitTaskPerDay, utils.GetCurrentDate()); err != nil {
			return nil, err
		}
	} else if res.Error != nil {
		return nil, res.Error
	}

	// Select to update
	res = db.Clauses(clause.Locking{
		Strength: "UPDATE",
		Options: "NOWAIT",
	}).Where(&models.UserConfigs{
		Email: email,
		Date: date,
	}).First(config)

	if res.Error != nil {
		return nil, res.Error
	}

	return config, nil
}

func SetConfig(email string, limit int8, date time.Time) (*models.UserConfigs, error) {
	config := &models.UserConfigs{
		Email: email,
		Limit: limit,
		Date: date,
	}
	// TODO: CreatedAt time is updated when conflict, need to fix
	res := db.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "Email"}, {Name: "Date"}},
		DoUpdates: clause.Assignments(map[string]interface{}{
			"Limit": limit,
		}),
	}).Create(&config)

	if res.Error != nil {
		log.WithError(res.Error).Error("Create User Config fail")
		return nil, res.Error
	}
	return config, nil
}
