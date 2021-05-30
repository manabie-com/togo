package postgre

import (
	"github.com/manabie-com/togo/db"
	"github.com/manabie-com/togo/modules/users"
)

type UsersRepo struct {
}

func (r UsersRepo) CheckLogin(userId string, pass string) (users.Users, error) {
	var users users.Users
	db.DB.Where("id = ? AND password = ?", userId, pass).Find(&users)
	return users, nil
}
