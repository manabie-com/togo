package service

import (
	"time"

	"github.com/go-chi/jwtauth/v5"
	c "github.com/manabie-com/togo/internal/pkg/config"
	d "github.com/manabie-com/togo/internal/todo/domain"
)

type AuthService struct {
	UserRepo d.UserRepository
}

func NewAuthService(userRepo d.UserRepository) *AuthService {
	return &AuthService{userRepo}
}

func (s *AuthService) ValidateUser(tokenAuth *jwtauth.JWTAuth, cred d.UserAuthParam) (string, error) {
	user, err := s.UserRepo.GetByCredentials(cred.Username, cred.Password)
	if err != nil {
		return "", err
	}

	if user == nil {
		return "", nil
	}

	expiredPeriod := c.GetEnvInt("JWT_EXPIRED_PERIOD")
	claims := map[string]interface{}{
		"userID": user.ID,
		"exp":    time.Now().Add(time.Minute * time.Duration(expiredPeriod)).Unix(),
	}

	_, token, err := tokenAuth.Encode(claims)
	if err != nil {
		return "", err
	}

	return token, nil
}
