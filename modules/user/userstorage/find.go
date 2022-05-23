package userstorage

import (
	"context"
	"github.com/japananh/togo/common"
	"github.com/japananh/togo/modules/user/usermodel"

	"gorm.io/gorm"
)

func (s *sqlStore) FindUser(
	_ context.Context,
	conditions map[string]interface{},
	moreInfo ...string,
) (*usermodel.User, error) {
	db := s.db.Table(usermodel.User{}.TableName())

	for i := range moreInfo {
		db = db.Preload(moreInfo[i])
	}

	var user usermodel.User

	if err := db.Where(conditions).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, common.ErrRecordNotFound
		}
		return nil, common.ErrDB(err)
	}

	return &user, nil
}
