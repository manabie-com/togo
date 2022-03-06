package todo

import (
	"context"

	"github.com/pkg/errors"
	"gorm.io/gorm"

	"github.com/khangjig/togo/codetype"
	"github.com/khangjig/togo/model"
)

func NewPG(getDB func(ctx context.Context) *gorm.DB) Repository {
	return pgRepository{getDB}
}

type pgRepository struct {
	getDB func(ctx context.Context) *gorm.DB
}

func (p pgRepository) Create(ctx context.Context, todo *model.Todo) error {
	return p.getDB(ctx).Create(todo).Error
}

func (p pgRepository) GetByID(ctx context.Context, id int64) (*model.Todo, error) {
	var data model.Todo

	err := p.getDB(ctx).Where("id = ?", id).First(&data).Error
	if err != nil {
		return nil, errors.Wrap(err, "get by id")
	}

	return &data, nil
}

func (p pgRepository) Update(ctx context.Context, todo *model.Todo) error {
	return p.getDB(ctx).Save(todo).Error
}

func (p pgRepository) DeleteByID(ctx context.Context, id int64, unscoped bool) error {
	db := p.getDB(ctx)

	if unscoped {
		db = db.Unscoped()
	}

	return db.Where("id = ?", id).Delete(&model.Todo{}).Error
}

func (p pgRepository) GetList(
	ctx context.Context,
	userID int64,
	conditions interface{},
	search string,
	order string,
	paginator codetype.Paginator,
) ([]model.Todo, int64, error) {
	var (
		total  int64
		offset int
		data   = make([]model.Todo, 0)
	)

	db := p.getDB(ctx).Model(&model.Todo{}).Where("user_id = ?", userID)

	if conditions != nil {
		db = db.Where(conditions)
	}

	if search != "" {
		db.Where("MATCH (title, content) AGAINST (? IN NATURAL LANGUAGE MODE)", search)
	}

	if order != "" {
		db = db.Order(order)
	}

	if paginator.Page != 1 {
		offset = paginator.Limit * (paginator.Page - 1)
	}

	if paginator.Limit != -1 {
		err := db.Count(&total).Error
		if err != nil {
			return nil, 0, err
		}
	}

	err := db.Limit(paginator.Limit).Offset(offset).Find(&data).Error
	if err != nil {
		return nil, 0, err
	}

	if paginator.Limit == -1 {
		total = int64(len(data))
	}

	return data, total, nil
}
