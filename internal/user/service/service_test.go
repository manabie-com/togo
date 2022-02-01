package service

import (
	"context"
	"fmt"
	"log"
	"testing"

	"github.com/bwmarrin/snowflake"

	"github.com/manabie-com/togo/pkg/errorx"

	repository2 "github.com/manabie-com/togo/internal/task/repository"
	"github.com/manabie-com/togo/internal/user/repository"

	"github.com/manabie-com/togo/core/config"
	"github.com/manabie-com/togo/pkg/database"
	"github.com/stretchr/testify/assert"
)

var userServiceTest UserService

const (
	testUser     = "testpostgres"
	testPassword = "testpostgres"
	testDbName   = "testpostgres"
	testHost     = "localhost"
	testPort     = 5430
)

func TestMain(m *testing.M) {
	db := database.New(config.DBConfig{Test_PostgresDB: &config.PostgresConfig{
		Host:     testHost,
		Port:     testPort,
		Username: testUser,
		Password: testPassword,
		Database: testDbName,
		SSLMode:  "disable",
	}})
	userRepo := repository.NewUserRepository(db.TestManabieDB)
	taskRepo := repository2.NewTaskRepository(db.TestManabieDB)
	userServiceTest = NewUserService(userRepo, taskRepo, db.TestManabieDB)
	m.Run()
	d, err := db.TestManabieDB.DB()
	if err != nil {
		log.Fatalf("Could not connect sql database: %s", err)
	}
	defer d.Close()
}

func Test_Login(t *testing.T) {
	t.Run("Test_Login_Success", func(t *testing.T) {
		userRes, err := userServiceTest.Login(context.Background(), &LoginUserArgs{
			Username: "example",
			Password: "123456789",
		})
		assert.Nil(t, err)
		assert.NotNil(t, userRes)
	})

	t.Run("Test_Login_Wrong_Password", func(t *testing.T) {
		_, err := userServiceTest.Login(context.Background(), &LoginUserArgs{
			Username: "example",
			Password: "example_123456789",
		})
		if assert.NotNil(t, err) {
			e := err.(errorx.ErrorInterface)
			assert.Equal(t, errorx.ErrAuthFailure(err).GetTitle(), e.GetTitle())
		}
	})
}

func Test_SaveUser(t *testing.T) {
	t.Run("Test_CreateUser_Success", func(t *testing.T) {
		randomUsername := GearedID()
		err := userServiceTest.CreateUser(context.Background(), &CreateUserArgs{
			Username:  randomUsername,
			Password:  "123456789",
			TaskLimit: 10,
		})
		assert.Nil(t, err)
	})
	t.Run("Test_SaveUser_Fail_Missing_Arguments", func(t *testing.T) {
		err := userServiceTest.CreateUser(context.Background(), &CreateUserArgs{})
		if assert.NotNil(t, err) {
			e := err.(errorx.ErrorInterface)
			assert.Equal(t, errorx.ErrInvalidParameter(err).GetTitle(), e.GetTitle())
		}
	})
}

func Test_GetUser(t *testing.T) {
	t.Run("Test_GetUser_Success", func(t *testing.T) {
		user, err := userServiceTest.GetUser(context.Background(), &GetUserArgs{
			UserID: 1,
		})
		assert.Nil(t, err)
		assert.Equal(t, 1, user.ID)
		assert.Equal(t, "$2a$14$QsCCoXE0O84KFzMMOfMQ8O.szZqS9pRzaxAovr4/WAN0tgO//A9S2", user.Password)
		assert.Equal(t, 10, user.TaskLimit)
	})
	t.Run("Test_SaveUser_Fail_Not_Found", func(t *testing.T) {
		user, err := userServiceTest.GetUser(context.Background(), &GetUserArgs{
			UserID: 10000000,
		})
		assert.Nil(t, user)
		if assert.NotNil(t, err) {
			e := err.(errorx.ErrorInterface)
			assert.Equal(t, errorx.ErrUserNotFound(err).GetTitle(), e.GetTitle())
		}
	})
}

func GearedID() string {
	n, err := snowflake.NewNode(1)
	if err != nil {
		return ""
	}
	return fmt.Sprint(n.Generate().Int64())
}
