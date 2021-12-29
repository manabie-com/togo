package repositories

import (
	"github.com/manabie-com/togo/entities"
	"gorm.io/gorm"
)

type IUserRepository interface {
	GetLimitTaskPerDay(userId uint64) (int, error)
}

type userConnection struct {
	connection *gorm.DB
}

func NewUserRepository(db *gorm.DB) IUserRepository {
	return &userConnection{
		connection: db,
	}
}

func (userConn *userConnection) GetLimitTaskPerDay(userId uint64) (int, error) {
	var user entities.User
	err := userConn.connection.Where("id = ?", userId).Find(&user).Error

	if err != nil {
		return 0, err
	}
	return int(user.LimitTask), nil
}
