package users

import "todo/database"

type UserService interface {
	GetAll(data interface{}) error
}

type userService struct {
	responstory database.Responstory
}

func InitUserService(responstory database.Responstory) UserService {
	service := userService{responstory: responstory}
	return service
}

func (s userService) GetAll(data interface{}) error {
	return nil
}
