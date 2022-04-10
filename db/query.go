package db

import (
	"github.com/namnhatdoan/togo/models"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"time"
)

func GetConfigForUpdate(email string, date time.Time) (*models.UserConfigs, error) {
	config := &models.UserConfigs{}
	res := db.Clauses(clause.Locking{
		Strength: "UPDATE",
		Options: "NOWAIT",
	}).Where(&models.UserConfigs{
		Email: email,
		Date: date,
	}).First(config)

	return config, res.Error
}

func UpsertUserConfig(email string, limit int8, date time.Time) (*models.UserConfigs, error) {
	config := &models.UserConfigs{
		Email: email,
		Limit: limit,
		Date: date,
	}
	res := db.Where(&models.UserConfigs{
		Email: email,
		Date: date,
	}).First(config)

	if res.Error == nil {
		// Config found > update
		config.Limit = limit
		if res := db.Save(config); res.Error != nil {
			log.WithField("email", email).
				WithField("date", date).
				WithField("limit", limit).
				WithError(res.Error).
				Error("Update User Config fail")
			return nil, res.Error
		}
	} else if res.Error == gorm.ErrRecordNotFound {
		// Config not found > create
		if res := db.Create(config); res.Error != nil {
			log.WithField("email", email).
				WithField("date", date).
				WithField("limit", limit).
				WithError(res.Error).
				Error("Create User Config fail")
			return nil, res.Error
		}
	} else {
		// DB Error > return
		return nil, res.Error
	}

	return config, nil
}
