package services

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"testing"

	_ "github.com/lib/pq"
	"github.com/manabie-com/togo/internal/config"
	postgres "github.com/manabie-com/togo/internal/storages/postgres"
	"github.com/stretchr/testify/assert"
)

var (
	todoSvc *ToDoService
	db      *sql.DB
	userID  = "akagi"
	token   string
)

func init() {
	var err error

	conf := config.GetConfig()
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		conf.PGHost, conf.PGPort, conf.PGUser, conf.PGPassword, conf.PGDBName)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal("error opening db", err)
	}
	todoSvc = &ToDoService{
		JWTKey: "",
		Store: &postgres.PostgresDB{
			DB: db,
		},
	}
}

func TestCreateToken(t *testing.T) {
	var err error
	assert := assert.New(t)

	token, err = todoSvc.createToken(userID)
	assert.Nil(err)
	assert.NotNil(token)
	// fmt.Printf("token: %v\n", token)
}

func TestValidToken(t *testing.T) {
	assert := assert.New(t)

	req := new(http.Request)
	req.Header = make(http.Header)
	req.Header.Set("Authorization", "Bearer "+token)
	var ok bool
	req, ok = todoSvc.validToken(req)
	assert.True(ok)
	assert.NotNil(req)
	// fmt.Printf("request: %+v\n", req)
}

func TestValue(t *testing.T) {
	assert := assert.New(t)

	req := new(http.Request)
	req.Form = make(url.Values)
	req.Form.Set("user_id", "firstUser")
	req.Form.Set("password", "example")

	userID := value(req, "user_id")
	password := value(req, "password")
	assert.NotNil(userID)
	assert.Equal("firstUser", userID.String)
	assert.Equal(true, userID.Valid)
	assert.NotNil(password)
	assert.Equal("example", password.String)
	assert.Equal(true, password.Valid)
}
