package task

import (
	"context"

	"gorm.io/gorm"
)

type task struct {
	userID int
	db     *gorm.DB
}

func Initalize(ctx context.Context) task {
	userID := ctx.Value("userID").(int)
	db := ctx.Value("gorm").(*gorm.DB)

	return task{
		userID: userID,
		db:     db,
	}
}
