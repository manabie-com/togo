package impl

import (
	"errors"
	"github.com/jinzhu/gorm"
	"github.com/manabie-com/togo/internal/model"
	"github.com/manabie-com/togo/internal/repository"
	"github.com/manabie-com/togo/internal/repository/impl"
	"github.com/manabie-com/togo/internal/utils"
	"net/http"
)

type UserServiceImpl struct{
	repository repository.UsersRepository
}

func NewUserServiceImpl(db *gorm.DB) *UserServiceImpl {
	return &UserServiceImpl{impl.NewUserRepositoryImpl(db)}
}

// ValidateUser returns tasks if match userID AND password
func (s *UserServiceImpl) ValidateUser(userID, pwd string) bool {
	var users model.Users
	users, err := s.repository.GetUserByIdAndOPassword(userID, pwd)
	if err != nil {
		return false
	}

	if users != (model.Users{}) {
		return true
	} else {
		return false
	}

}

func (s *UserServiceImpl) Login(userID, pwd string) (token string, err error, code int) {
	if s.ValidateUser(userID, pwd) {
		jwt := utils.NewJwt()
		token, err := jwt.CreateToken(userID)
		if err != nil {
			return "", err, http.StatusInternalServerError
		}
		return token, nil, http.StatusOK
	} else {
		return "", errors.New("incorrect user_id/pwd"), http.StatusUnauthorized
	}
}
