package service

import (
	"github.com/ansidev/togo/auth/dto"
	"github.com/ansidev/togo/domain/auth"
	"github.com/ansidev/togo/domain/user"
	"github.com/ansidev/togo/errs"
	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
)

type IAuthService interface {
	Login(usernamePasswordCredential dto.UsernamePasswordCredential) (string, error)
	GetCredential(token string) (dto.UserCredential, error)
}

func NewAuthService(userRepository user.IUserRepository, credRepository auth.ICredRepository) IAuthService {
	return &AuthService{userRepository, credRepository}
}

type AuthService struct {
	userRepository user.IUserRepository
	credRepository auth.ICredRepository
}

func (s *AuthService) Login(usernamePasswordCredential dto.UsernamePasswordCredential) (string, error) {
	userModel, err := s.userRepository.FindFirstByUsername(usernamePasswordCredential.Username)

	if err != nil {
		return "", errs.New(errs.ErrUsernameNotFound).
			WithCode(errs.ErrCodeUsernameNotFound).
			WithErr(err).
			Build()
	}

	hashedPassword := []byte(userModel.Password)
	requestPassword := []byte(usernamePasswordCredential.Password)
	err1 := bcrypt.CompareHashAndPassword(hashedPassword, requestPassword)

	if err1 != nil {
		return "", errs.New(errs.ErrWrongUserPassword).
			WithCode(errs.ErrCodeWrongUserPassword).
			WithErr(err1).
			Build()
	}

	token, err2 := s.credRepository.Save(userModel)

	if err2 != nil {
		return "", errs.New(errs.ErrCouldNotSaveToken).
			WithCode(errs.ErrCodeCouldNotSaveToken).
			WithErr(err2).
			Build()
	}

	return token, nil
}

func (s *AuthService) GetCredential(token string) (dto.UserCredential, error) {
	credential, err := s.credRepository.Get(token)

	if err != nil {
		e := errors.Cause(err)

		switch e {
		case errs.ErrRecordNotFound:
			return dto.UserCredential{},
				errs.New(errs.ErrTokenNotFound).
					WithCode(errs.ErrCodeTokenNotFound).
					WithErr(e).
					Build()
		case errs.ErrInternalAppFailure:
			return dto.UserCredential{},
				errs.New(errs.ErrInternalAppError).
					WithCode(errs.ErrCodeInternalAppError).
					WithErr(e).
					Build()
		default:
			return dto.UserCredential{},
				errs.New(errs.ErrInternalServiceError).
					WithCode(errs.ErrCodeRedisError).
					WithErr(e).
					Build()
		}
	}

	return dto.UserCredential{
		ID:           credential.ID,
		MaxDailyTask: credential.MaxDailyTask,
	}, nil
}
