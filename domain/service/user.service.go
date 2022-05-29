package service

import (
	"context"
	"golang.org/x/crypto/bcrypt"
	"math/rand"
	"togo/domain/errdef"
	"togo/domain/model"
	"togo/domain/repository"
)

type UserService interface {
	Register(ctx context.Context, user model.User) error
	Login(ctx context.Context, username string, password string) (model.User, error)
}

type userService struct {
	db           repository.UserRepository
	tokenService TokenService
}

func (this userService) Register(ctx context.Context, user model.User) error {
	if user.Limit == 0 {
		user.Limit = rand.Intn(3) + 2
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	user.Password = string(hashedPassword)
	err = this.db.Create(ctx, user)
	return err
}

func (this *userService) createToken(ctx context.Context, user *model.User) (string, error) {
	// TODO: create JWT token and saving it in database
	return "token", nil
}

func (this userService) Login(ctx context.Context, username string, password string) (model.User, error) {
	user, err := this.db.Get(ctx, username)
	if err != nil {
		return model.User{}, err
	}
	if err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return model.User{}, errdef.InvalidUsernameOrPassword
	}
	// create token for this session
	token, err := this.tokenService.CreateToken(ctx, user)
	user.Token = token
	return user, nil
}

func NewUserService(userRepo repository.UserRepository, tokenService TokenService) UserService {
	return &userService{db: userRepo, tokenService: tokenService}
}
