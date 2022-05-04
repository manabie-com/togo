package factories

import (
	"errors"

	"github.com/manabie-com/togo/handlers"

	"github.com/manabie-com/togo/models"
)

func UserFactory(typeAction string, togo *models.Togo) (*models.User, error) {
	if typeAction == "get" {
		return handlers.GetUserById(togo)
	}
	if typeAction == "add" {
		return handlers.CreateUser(togo)
	}
	return &models.User{}, errors.New("incorrect type action")
}
