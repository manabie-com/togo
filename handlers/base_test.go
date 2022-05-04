package handlers

import (
	"github.com/manabie-com/togo/database"
	"github.com/manabie-com/togo/models"
)

func SetUpUnitTest(todo *models.Togo) {
	database.ConnectDatabase()
}

func cleanUnitTest(todo *models.Togo) {
	database.DisconnectDatabase()
}

func cleanLimitTask(todo *models.Togo) {
	resetLimitTask(todo)
}
