package postgresql

import (
	"github.com/gin-gonic/gin"
	"github.com/shanenoi/togo/internal/storages/models"
)

type UserPostgreSQL interface {
	CreateUser(task *models.User) error
	FindUserByUserName(username string) (*models.User, error)
	FindUserById(uid string) (*models.User, error)
}

type userPostgreSQL struct {
	ctx *gin.Context
}

func NewUserPostgreSQL(ctx *gin.Context) UserPostgreSQL {
	return &userPostgreSQL{ctx}
}

func (tpsql *userPostgreSQL) CreateUser(user *models.User) error {
	db, _ := GetDb(tpsql.ctx)
	return db.Database.Create(user).Error
}
func (tpsql *userPostgreSQL) FindUserByUserName(username string) (*models.User, error) {
	db, _ := GetDb(tpsql.ctx)
	user := &models.User{}
	err := db.Database.Where("username = ?", username).First(user).Error
	return user, err
}
func (tpsql *userPostgreSQL) FindUserById(uid string) (*models.User, error) {
	db, _ := GetDb(tpsql.ctx)
	user := &models.User{}
	err := db.Database.Where("id = ?", uid).First(user).Error
	return user, err
}
