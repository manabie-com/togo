package dto

import (
	"github.com/trinhdaiphuc/togo/database/ent"
	"github.com/trinhdaiphuc/togo/internal/entities"
)

func User2UserEntity(user *ent.User) *entities.User {
	return &entities.User{
		ID:        user.ID,
		Username:  user.Username,
		Password:  user.Password,
		TaskLimit: user.TaskLimit,
	}
}
