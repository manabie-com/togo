package user

import (
	"context"

	"github.com/pkg/errors"
	"gorm.io/gorm"

	"github.com/khangjig/togo/model"
)

func NewPG(getDB func(ctx context.Context) *gorm.DB) Repository {
	return pgRepository{getDB}
}

type pgRepository struct {
	getDB func(ctx context.Context) *gorm.DB
}

func (p pgRepository) GetByID(ctx context.Context, id int64) (*model.User, error) {
	var data model.User

	err := p.getDB(ctx).Where("id = ?", id).First(&data).Error
	if err != nil {
		return nil, errors.Wrap(err, "get by id")
	}

	return &data, nil
}

func (p pgRepository) GetByEmail(ctx context.Context, email string) (*model.User, error) {
	var data model.User

	err := p.getDB(ctx).Where("email = ?", email).First(&data).Error
	if err != nil {
		return nil, errors.Wrap(err, "get by email")
	}

	return &data, nil
}
