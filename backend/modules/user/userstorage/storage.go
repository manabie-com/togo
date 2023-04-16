package userstorage

import (
	"context"

	"github.com/phathdt/libs/go-sdk/sdkcm"
	"gorm.io/gorm"
	"togo/modules/user/usermodel"
)

type sqlStore struct {
	db *gorm.DB
}

func NewSQLStore(db *gorm.DB) *sqlStore {
	return &sqlStore{db: db}
}

func (s *sqlStore) FindUser(ctx context.Context, conditions map[string]interface{}) (*usermodel.User, error) {
	db := s.db.Table(usermodel.User{}.TableName())

	var user usermodel.User

	if err := db.Where(conditions).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, sdkcm.RecordNotFound
		}

		return nil, sdkcm.ErrDB(err)
	}

	return &user, nil
}

func (s *sqlStore) CreateUser(ctx context.Context, data *usermodel.UserCreate) error {
	if err := s.db.Table(data.TableName()).Create(data).Error; err != nil {
		return sdkcm.ErrDB(err)
	}

	return nil
}
