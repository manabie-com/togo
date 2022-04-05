package user

import (
	"context"

	"gorm.io/gorm"
)

type user struct {
	db       *gorm.DB
	username string
}

func Initialize(ctx context.Context) user {
	db := ctx.Value("db").(*gorm.DB)
	username := ctx.Value("username").(string)

	return user{
		db:       db,
		username: username,
	}
}
