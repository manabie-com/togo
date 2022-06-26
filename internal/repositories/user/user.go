package user

import (
	"github.com/manabie-com/togo/internal/models"
	"github.com/manabie-com/togo/internal/pkg/id"
	"github.com/manabie-com/togo/internal/usecases/user"

	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
)

type repository struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) user.UserRepository {
	return &repository{
		DB: db,
	}
}

func New(user *models.User) *models.User {
	return &models.User{
		ID:            id.NewID().String(),
		Username:      user.Username,
		Password:      user.Password,
		MaxTaskPerDay: user.MaxTaskPerDay,
	}
}

// Create implements user.UserRepository
func (r *repository) Create(inputUser *models.User) error {
	if inputUser == nil || inputUser.Username == "" || inputUser.Password == "" {
		return errors.New("Input valid")
	}

	if err := r.DB.Create(inputUser).Error; err != nil {
		return errors.Wrap(err, "Fail create user")
	}

	return nil
}

// Login implements user.UserRepository
func (r *repository) Login(username, password string) (*models.User, error) {
	if username == "" || password == "" {
		return nil, errors.New("Input empty")
	}

	user := &models.User{}
	if err := r.DB.Where("username = ? and password = ?", username, password).Take(user).Error; err != nil {
		return nil, errors.Wrap(err, "Fail query user")
	}

	if user == nil {
		return nil, errors.New("User is not exists")
	}

	return user, nil
}
