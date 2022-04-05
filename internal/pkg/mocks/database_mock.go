package mocks

import (
	sqlmock "github.com/DATA-DOG/go-sqlmock"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// NewDatabaseMock func
func NewDatabaseMock() (sqlmock.Sqlmock, *gorm.DB) {
	mockDB, sqlMock, _ := sqlmock.New()
	gormDB, _ := gorm.Open(postgres.New(postgres.Config{
		Conn: mockDB,
	}), &gorm.Config{})
	return sqlMock, gormDB
}
