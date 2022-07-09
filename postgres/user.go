package postgres

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/lawtrann/togo"
)

type UserRepo struct {
	DB *DB
}

func NewUserRepo(db *DB) *UserRepo {
	return &UserRepo{DB: db}
}

func (ur *UserRepo) GetUserByName(ctx context.Context, username string) (*togo.User, error) {
	var result togo.User

	// Retrieve user by name
	fmt.Println(ur.DB.Now().Format("2006-01-02 15:04:05"), "\n", ISQLTemplate("UserRepoGetUserByName.sql"))
	err := ur.DB.DB.QueryRow(ISQLTemplate("UserRepoGetUserByName.sql"),
		username).Scan(&result.ID, &result.Username, &result.LimitedPerDay)
	if err != nil && err != sql.ErrNoRows {
		fmt.Println(err)
		return &togo.User{}, err
	}

	return &result, nil
}

func (ur *UserRepo) IsExceedPerDay(ctx context.Context, u *togo.User) (bool, error) {
	result := false

	// Check if exceed per day
	fmt.Println(ur.DB.Now().Format("2006-01-02 15:04:05"), "\n", ISQLTemplate("UserRepoIsExceedLimitedPerDay.sql"))
	err := ur.DB.DB.QueryRow(ISQLTemplate("UserRepoIsExceedLimitedPerDay.sql"),
		u.ID).Scan(&result)
	if err != nil && err != sql.ErrNoRows {
		fmt.Println(err)
		return false, err
	}

	return result, nil
}
