package handlers

import (
	"github.com/manabie-com/togo/database"
	"github.com/manabie-com/togo/models"
)

func SetUpUnitTest(togo *models.Togo) {
	database.ConnectDatabase()
}

func cleanUnitTest(togo *models.Togo) {
	database.DisconnectDatabase()
}

func cleanLimitTask(togo *models.Togo) {
	resetLimitTask(togo)
}
