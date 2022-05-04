package factories

import (
	"errors"

	"github.com/manabie-com/togo/handlers"
	"github.com/manabie-com/togo/models"
)

func TodoFactory(typeAction string, todo *models.Togo) (*models.User, error) {
	if typeAction == "add" {
		return handlers.AddTodo(todo)
	}
	return nil, errors.New("incorrect type action")
}
