package repository

import (
	"ansidev.xyz/pkg/log"
	"github.com/ansidev/togo/domain/user"
	"github.com/ansidev/togo/errs"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"strconv"
)

func NewPostgresUserRepository(db *gorm.DB) user.IUserRepository {
	return &postgresUserRepository{db}
}

type postgresUserRepository struct {
	db *gorm.DB
}

func (r *postgresUserRepository) FindFirstByUsername(username string) (user.User, error) {
	var a user.User
	result := r.db.First(&a, "username = ?", username)

	if result.Error != nil {
		err := result.Error

		if errors.Is(err, gorm.ErrRecordNotFound) {
			return user.User{}, errors.Wrap(errs.ErrRecordNotFound, result.Error.Error())
		}

		log.Errorz("Error while querying user by username = "+username, zap.Error(result.Error))
		return user.User{}, errors.Wrap(errs.ErrDatabaseFailure, result.Error.Error())
	}

	return a, nil
}

func (r *postgresUserRepository) FindFirstByID(id int64) (user.User, error) {
	var userModel user.User
	result := r.db.First(&userModel, id)

	if result.Error != nil {
		err := result.Error

		if errors.Is(err, gorm.ErrRecordNotFound) {
			return user.User{}, errors.Wrap(errs.ErrRecordNotFound, result.Error.Error())
		}

		log.Errorz("Error while querying user by id = "+strconv.FormatInt(id, 10), zap.Error(result.Error))
		return user.User{}, errors.Wrap(errs.ErrDatabaseFailure, result.Error.Error())
	}

	return userModel, nil
}
