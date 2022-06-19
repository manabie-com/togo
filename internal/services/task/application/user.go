package application

import (
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"togo/internal/services/task/domain"
	"togo/internal/services/task/store/postgres"
)

type UserService struct {
	userRepo *postgres.UserRepository
}

func NewUserService(userRepo *postgres.UserRepository) *UserService {
	return &UserService{userRepo}
}

type AddUserCommand struct {
	DailyTaskLimit int
}

func (s *UserService) CreateUser(cmd AddUserCommand) (*domain.User, error) {
	entity := domain.NewUser(uuid.New(), cmd.DailyTaskLimit)
	if err := s.userRepo.Save(entity); err != nil {
		return nil, errors.Wrap(err, "cannot save user")
	}
	return entity, nil
}

func (s *UserService) FindUserById(id uuid.UUID) (*domain.User, error) {
	entity, err := s.userRepo.FindByID(id)
	if err != nil {
		return nil, errors.Wrapf(err, "cannot get user %s", id)
	}
	return entity, nil
}
