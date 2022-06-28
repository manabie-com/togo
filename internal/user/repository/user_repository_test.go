package repository

import (
	"errors"
	"github.com/stretchr/testify/require"
	"gopkg.in/DATA-DOG/go-sqlmock.v1"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"testing"
	"time"
)

func TestUserRepository_IsUserExisted_Existed(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	gormDB, err := gorm.Open(postgres.New(postgres.Config{
		Conn: db,
	}), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
		NamingStrategy: schema.NamingStrategy{
			SingularTable: false,
		},
	})
	require.NoError(t, err)
	userId := 1
	mock.ExpectPrepare("SELECT * FROM \"users\" WHERE id = $1 AND \"users\".\"deleted_at\" IS NULL LIMIT 1").
		ExpectQuery().WithArgs(userId).WillReturnError(nil)
	userRepo := NewUserRepository(gormDB)
	err = userRepo.IsUserExisted(int64(userId))
	require.NoError(t, err)
}

func TestUserRepository_IsUserExisted_NotExisted(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	gormDB, err := gorm.Open(postgres.New(postgres.Config{
		Conn: db,
	}), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
		NamingStrategy: schema.NamingStrategy{
			SingularTable: false,
		},
	})
	require.NoError(t, err)
	userId := 1
	mock.ExpectPrepare("SELECT * FROM \"users\" WHERE id = $1 AND \"users\".\"deleted_at\" IS NULL LIMIT 1").
		ExpectQuery().WithArgs(userId).WillReturnError(errors.New("some error"))
	userRepo := NewUserRepository(gormDB)
	err = userRepo.IsUserExisted(int64(userId))
	require.Error(t, err)
	require.Equal(t, errors.New("user is not existed"), err)
}

func TestUserRepository_IsUserHavingMaxTodo_HavingMaxTodo(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	gormDB, err := gorm.Open(postgres.New(postgres.Config{
		Conn: db,
	}), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
		NamingStrategy: schema.NamingStrategy{
			SingularTable: false,
		},
	})
	require.NoError(t, err)
	userId := 1
	date := time.Now()
	mock.ExpectPrepare("SELECT \\* FROM \"users\" WHERE (id = $1 AND max_todo > (SELECT COUNT(\\*) FROM \"todos\" WHERE CAST(created_at as DATE) = $2 AND user_id = $1)) AND \"users\".\"deleted_at\" IS NULL LIMIT 1").
		ExpectQuery().WithArgs(userId, date.Format(DateFormat)).WillReturnError(errors.New("record not found"))
	userRepo := NewUserRepository(gormDB)
	err = userRepo.IsUserHavingMaxTodo(int64(userId), date)
	require.Error(t, err)
	require.Equal(t, errors.New("user have too many todos"), err)
}

func TestUserRepository_IsUserHavingMaxTodo_NotHavingMaxTodo(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	gormDB, err := gorm.Open(postgres.New(postgres.Config{
		Conn: db,
	}), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
		NamingStrategy: schema.NamingStrategy{
			SingularTable: false,
		},
	})
	require.NoError(t, err)
	userId := uint(1)
	date := time.Now()
	mock.ExpectPrepare(`SELECT.*`).
		ExpectQuery().WithArgs(userId, date.Format(DateFormat)).WillReturnError(nil)
	userRepo := NewUserRepository(gormDB)
	err = userRepo.IsUserHavingMaxTodo(int64(userId), date)
	require.NoError(t, err)
}
