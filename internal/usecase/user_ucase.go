package usecase

import (
	"context"
	"database/sql"

	"github.com/manabie-com/togo/internal/storages"
	"github.com/manabie-com/togo/internal/storages/postgres"
)

// UserUsecase implement usecase of user
type UserUsecase struct {
	Store *postgres.PostgresDB
}

// GetUser get user info
func (u *UserUsecase) GetUser(ctx context.Context, userID sql.NullString) (*storages.User, error) {
	return u.Store.GetUser(ctx, userID)
}

// ValidateUser match userID, password
func (u *UserUsecase) ValidateUser(ctx context.Context, userID, pwd sql.NullString) bool {
	return u.Store.ValidateUser(ctx, userID, pwd)
}
