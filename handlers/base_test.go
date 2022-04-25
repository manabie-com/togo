package handlers

import (
	"github.com/manabie-com/togo/database"
	"github.com/manabie-com/togo/models"
)

func SetUpUnitTest(todo *models.Todo) {
	database.ConnectDatabase()
}

func cleanUnitTest(todo *models.Todo) {
	database.DisconnectDatabase()
}

func cleanLimitTask(todo *models.Todo) {
	resetLimitTask(todo)
}
