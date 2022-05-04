package factories

import (
	"errors"

	"github.com/manabie-com/togo/handlers"
	"github.com/manabie-com/togo/models"
)

func TogoFactory(typeAction string, togo *models.Togo) (*models.User, error) {
	if typeAction == "add" {
		return handlers.Addtogo(togo)
	}
	return nil, errors.New("incorrect type action")
}
