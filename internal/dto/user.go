package dto

import (
	"todo/database/ent"
	"todo/internal/entities"
)

func User2UserEntity(user *ent.User) *entities.User {
	return &entities.User{
		ID:        user.ID,
		Username:  user.Username,
		Password:  user.Password,
		TaskLimit: user.TaskLimit,
	}
}
