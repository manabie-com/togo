package users

import (
	"github.com/manabie-com/togo/middlewares"
	"github.com/manabie-com/togo/modules/common"
)

type Service struct {
	Repo UsersRepository
}

func (r Service) CheckLogin(userId string, pass string) (Users, *common.RequestError) {

	user, err := r.Repo.CheckLogin(userId, pass)

	if err != nil {
		return user, common.NewInternalError()
	}
	return user, nil

}

func (r Service) Login(userId string, pass string) (string, *common.RequestError) {

	user, err := r.CheckLogin(userId, pass)
	if err != nil {
		return "", err
	}

	if len(user.ID) < 1 {
		return "", common.NewUnAuthorizeError("userId or password is incorrect!")
	}

	token, errToken := middlewares.CreateToken(user.ID, user.MaxTodo)
	if errToken != nil {
		return "", common.NewCustomInternalError("Error generate token!")
	}

	return token, nil
}
