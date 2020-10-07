package user

import (
	"errors"

	"github.com/manabie-com/togo/internal/dto"
	"github.com/manabie-com/togo/internal/util"
)

// Service interface
type Service interface {
	GetAll() ([]User, error)
	GetUser(id uint64) (User, error)
	Login(dto *dto.LoginDTO) (User, error)
	CheckEmailExist(email string) (bool, error)
	Create(email string, password string, maxToto int) (User, error)
}

// NewUserService func
func NewUserService(repository Repository) (Service, error) {
	return &service{
		repository: repository,
	}, nil
}

type service struct {
	repository Repository
}

func (service *service) GetAll() ([]User, error) {
	return service.repository.GetAll()
}
func (service *service) GetUser(id uint64) (User, error) {
	return service.repository.GetUser(id)
}

func (service *service) Login(dto *dto.LoginDTO) (User, error) {
	user, err := service.repository.GetUserByEmail(dto.Email)
	if err != nil {
		return User{}, errors.New("wrong email")
	}
	err = util.CompareHashPassword(dto.Password, user.Password)
	if err != nil {
		return User{}, errors.New("wrong password")
	}
	return user, nil
}

func (service *service) CheckEmailExist(email string) (bool, error) {
	return service.repository.CheckEmailExist(email)
}

func (service *service) Create(email string, password string,maxToto int) (User, error) {
	return service.repository.Create(email, password, maxToto)
}
