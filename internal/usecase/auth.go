package usecase

import (
	"errors"
	"fmt"

	"github.com/manabie-com/togo/internal/domain"
	"golang.org/x/crypto/bcrypt"
)

type AuthUCConf struct {
	Hash    HashConf
	Default struct {
		MaxTaskPerDay int
	}
}

type HashConf struct {
	Secret string
	Algo   string
	Cost   int
}

type userUseCase struct {
	userStore         domain.UserStore
	config            AuthUCConf
	passwordValidator func([]byte, []byte) bool
	passwordGenerator func([]byte) ([]byte, error)
}

func NewAuthUseCase(c AuthUCConf, userStore domain.UserStore) (userUseCase, error) {
	result := userUseCase{
		config:    c,
		userStore: userStore,
	}
	switch c.Hash.Algo {
	case "bcrypt":
		result.passwordValidator = bcryptValidator
		result.passwordGenerator = bcryptPasswordGenerator(c.Hash.Cost)
	//TODO: more algorithm supported
	default:
		return userUseCase{}, fmt.Errorf("unsupported password hashing alorithm: %s", c.Hash.Algo)
	}
	return result, nil
}

func (u userUseCase) ValidateUser(userID, rawPassword string) (bool, error) {
	usr, err := u.userStore.FindUserByID(userID)
	if err != nil {
		if errors.Is(err, domain.UserNotFound(userID)) {
			return false, nil
		}
		return false, err
	}
	return u.passwordValidator([]byte(rawPassword), []byte(usr.Password)), nil
}

func (u userUseCase) CreateUser(id, pswd string) error {
	hashed, err := u.passwordGenerator([]byte(pswd))
	if err != nil {
		return err
	}
	return u.userStore.CreateUser(domain.User{
		ID:             id,
		Password:       string(hashed),
		MaxTasksPerDay: u.config.Default.MaxTaskPerDay,
	})
}

func bcryptValidator(given, hashed []byte) bool {
	err := bcrypt.CompareHashAndPassword(hashed, given)
	return err == nil
}

func bcryptPasswordGenerator(cost int) func([]byte) ([]byte, error) {
	return func(raw []byte) ([]byte, error) {
		bs, err := bcrypt.GenerateFromPassword(raw, cost)
		if err != nil {
			return nil, fmt.Errorf("failed to generate password: %s", err)
		}
		return bs, nil
	}
}
