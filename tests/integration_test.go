package tests

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"testing"

	"github.com/bwmarrin/snowflake"
	"github.com/stretchr/testify/assert"

	"github.com/manabie-com/togo/pkg/database"
	"github.com/pressly/goose/v3"

	"github.com/manabie-com/togo/internal/task/handler"
	userHandler "github.com/manabie-com/togo/internal/user/handler"

	"github.com/manabie-com/togo/core/config"
	"github.com/manabie-com/togo/registry"
)

const (
	testUser     = "testpostgres"
	testPassword = "testpostgres"
	testDbName   = "testpostgres"
	testHost     = "localhost"
	testPort     = 5430
)

var (
	integrationTest = &IntegrationTest{}
	db              = &sql.DB{}
)

type IntegrationTest struct {
	registry    *registry.Registry
	TaskHander  *handler.TaskHandler
	UserHandler *userHandler.UserHandler
}

func TestMain(m *testing.M) {
	var r *registry.Registry
	r, err := registry.New(config.Config{Databases: config.DBConfig{Test_PostgresDB: &config.PostgresConfig{
		Host:     testHost,
		Port:     testPort,
		Username: testUser,
		Password: testPassword,
		Database: testDbName,
		SSLMode:  "disable",
	}}})
	if err != nil {
		log.Fatalf("Could not registry: %s", err)
	}
	db, err = r.DB.TestManabieDB.DB()
	if err != nil {
		log.Fatalf("Could not registry: %s", err)
	}
	if err = db.Ping(); err != nil {
		log.Fatalf("Could not ping database: %s", err)
	}
	defer db.Close()
	integrationTest = &IntegrationTest{
		registry:    r,
		TaskHander:  handler.New(r.RegisterTaskService()),
		UserHandler: userHandler.New(r.RegisterUserService()),
	}
	os.Exit(m.Run())
}

func GearedID() string {
	n, err := snowflake.NewNode(1)
	if err != nil {
		return ""
	}
	return fmt.Sprint(n.Generate().Int64())
}

// CreateUser -> Login -> AddTask
func TestToDoService_AddTask(t *testing.T) {
	t.Run("Test_CreateUser", func(t *testing.T) {
		randomUsername := GearedID()
		createUserBody, err := json.Marshal(map[string]interface{}{
			"username":   randomUsername,
			"password":   "123456789",
			"task_limit": 5,
		})
		assert.Nil(t, err)
		createUserRes, err := http.Post("http://localhost:8080/api/user", "application/json", bytes.NewReader(createUserBody))
		assert.Nil(t, err)
		assert.Equal(t, http.StatusCreated, createUserRes.StatusCode)

		t.Run("Test_Login", func(t *testing.T) {
			loginReq, err := json.Marshal(map[string]string{
				"username": randomUsername,
				"password": "123456789",
			})
			assert.Nil(t, err)

			loginRes, err := http.Post("http://localhost:8080/api/user/login", "application/json", bytes.NewReader(loginReq))
			assert.Equal(t, http.StatusOK, loginRes.StatusCode)

			loginResStr, err := ioutil.ReadAll(loginRes.Body)
			assert.Nil(t, err)

			var loginResponse userHandler.LoginUserResponse

			assert.Nil(t, json.Unmarshal(loginResStr, &loginResponse))

			assert.NotNil(t, loginResponse)

			t.Run("Test_CreateTask", func(t *testing.T) {
				createTaskReq, err := json.Marshal(map[string]string{
					"content": "example content",
				})
				assert.Nil(t, err)

				req, err := http.NewRequest(http.MethodPost, "http://localhost:8080/api/task", bytes.NewReader(createTaskReq))
				assert.Nil(t, err)
				req.Header.Set("Content-Type", "application/json")
				req.Header.Set("Authorization", "Bearer "+loginResponse.AccessToken)
				createTaskRes, err := http.DefaultClient.Do(req)
				assert.Equal(t, http.StatusCreated, createTaskRes.StatusCode)
			})
		})

	})
}

func MigrateTestDB(database *database.Database) error {
	var err error
	if err = goose.SetDialect("postgres"); err != nil {
		return err
	}
	d, err := database.ManabieDB.DB()
	return goose.Up(d, "./../script/test_database_script")
}
