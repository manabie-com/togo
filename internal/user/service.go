package user

import "github.com/xrexonx/togo/internal/repository"

func GetById(id string) User {
	var user, _ = repository.GetByID[User](id)
	return user
}
