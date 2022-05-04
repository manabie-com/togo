package factories

import (
	"errors"

	"github.com/manabie-com/togo/handlers"

	"github.com/manabie-com/togo/models"
)

func UserFactory(typeAction string, todo *models.Togo) (*models.User, error) {
	if typeAction == "get" {
		return handlers.GetUserById(todo)
	}
	if typeAction == "add" {
		return handlers.CreateUser(todo)
	}
	return &models.User{}, errors.New("incorrect type action")
}
