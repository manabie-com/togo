package repository

import (
	"context"
	"fmt"
	"log"
	"testing"

	"github.com/bwmarrin/snowflake"

	"github.com/manabie-com/togo/pkg/errorx"

	"github.com/manabie-com/togo/model"

	"github.com/stretchr/testify/assert"

	"github.com/manabie-com/togo/core/config"
	"github.com/manabie-com/togo/pkg/database"
)

var (
	db = &database.Database{}
)
var userRepo UserRepository

const (
	testUser     = "testpostgres"
	testPassword = "testpostgres"
	testDbName   = "testpostgres"
	testHost     = "localhost"
	testPort     = 5430
)

func TestMain(m *testing.M) {
	db = database.New(config.DBConfig{Test_PostgresDB: &config.PostgresConfig{
		Host:     testHost,
		Port:     testPort,
		Username: testUser,
		Password: testPassword,
		Database: testDbName,
		SSLMode:  "disable",
	}})
	userRepo = NewUserRepository(db.TestManabieDB)
	m.Run()
	d, err := db.TestManabieDB.DB()
	if err != nil {
		log.Fatalf("Could not connect sql database: %s", err)
	}
	defer d.Close()
}

func Test_GetUser(t *testing.T) {
	t.Run("Test_GetUserByID", func(t *testing.T) {
		user, err := userRepo.GetUser(context.Background(), &model.User{
			ID: 1,
		})
		if err != nil {
			assert.Error(t, err)
		}
		assert.Equal(t, 1, user.ID)
		assert.Equal(t, "$2a$14$QsCCoXE0O84KFzMMOfMQ8O.szZqS9pRzaxAovr4/WAN0tgO//A9S2", user.Password)
		assert.Equal(t, 10, user.TaskLimit)
	})

	t.Run("Test_GetUserByID_Not_Found_Error", func(t *testing.T) {
		_, err := userRepo.GetUser(context.Background(), &model.User{
			ID: 10000,
		})
		if assert.NotNil(t, err) {
			e := err.(errorx.ErrorInterface)
			assert.Equal(t, errorx.ErrUserNotFound(err).GetTitle(), e.GetTitle())
		}
	})
}

func Test_SaveUser(t *testing.T) {
	t.Run("Test_SaveUser", func(t *testing.T) {
		randomUsername := GearedID()
		err := userRepo.SaveUser(db.TestManabieDB, &model.User{
			Username:  randomUsername,
			Password:  "123456789",
			TaskLimit: 5,
		})
		assert.Nil(t, err)
	})
}

func GearedID() string {
	n, err := snowflake.NewNode(1)
	if err != nil {
		return ""
	}
	return fmt.Sprint(n.Generate().Int64())
}
