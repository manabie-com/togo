package users

import (
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
)

func setupMock() (*gorm.DB, sqlmock.Sqlmock, error) {
	sqlDB, mock, err := sqlmock.New() // mock sql.DB
	if err != nil {
		return nil, nil, errors.Wrap(err, "Fail connection to database")
	}

	db, err := gorm.Open("postgres", sqlDB) // open gorm db
	if err != nil {
		return nil, nil, err
	}

	if db == nil {
		return nil, nil, errors.New("Fail connection to database")
	}

	return db, mock, nil
}

func SetUpRouter() *gin.Engine {
	router := gin.Default()
	return router
}
