package repository

import (
	entity "github.com/manabie-com/togo/internal/entities"
	postgre "github.com/manabie-com/togo/internal/storages/postgre"
)

// UserRepository action CRUD with Users entity
type UserRepository struct {
	Store *postgre.Storage
}

// ValidateUser func validate user
func (repo *UserRepository) ValidateUser(username string, password string) (*entity.User, error) {

	result, err := repo.Store.ValidateUser(username, password)

	return result, err
}
