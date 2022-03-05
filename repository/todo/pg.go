package todo

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

func (p pgRepository) Create(ctx context.Context, channel *model.Todo) error {
	return p.getDB(ctx).Create(channel).Error
}

func (p pgRepository) GetByID(ctx context.Context, id int64) (*model.Todo, error) {
	var data model.Todo

	err := p.getDB(ctx).Where("id = ?", id).First(&data).Error
	if err != nil {
		return nil, errors.Wrap(err, "get by id")
	}

	return &data, nil
}

func (p pgRepository) Update(ctx context.Context, channel *model.Todo) error {
	return p.getDB(ctx).Save(channel).Error
}

func (p pgRepository) Delete(ctx context.Context, channel *model.Todo) error {
	return p.getDB(ctx).Delete(channel).Error
}
