package user

import "github.com/xrexonx/togo/internal/repository"

// FindByID get user by ID
func FindByID(id string) User {
	var user, _ = repository.GetByID[User](id)
	return user
}
