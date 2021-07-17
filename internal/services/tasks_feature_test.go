package services

import (
	"context"
	"database/sql"
	"github.com/bxcodec/faker/v3"
	"github.com/gavv/httpexpect/v2"
	"github.com/manabie-com/togo/internal/helpers"
	"github.com/manabie-com/togo/internal/storages"
	sqllite "github.com/manabie-com/togo/internal/storages/sqlite"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

/**
* What is Feature Test?
* Basically, HTTP Test to test the endpoint if they are working properly.
 */

func beforeTest(t *testing.T) (*httpexpect.Expect, *httptest.Server, *sqllite.LiteDB) {
	db, err := sql.Open("sqlite3", "../../__fixtures__/data.db")
	if err != nil {
		log.Fatal("error opening db", err)
	}

	// create http.Handler
	liteDB := &sqllite.LiteDB{
		DB: db,
	}
	handler := &ToDoService{
		JWTKey: "wqGyEBBfPK9w3Lxw",
		Store:  liteDB,
	}

	server := httptest.NewServer(handler.ServeHTTP())
	return httpexpect.New(t, server.URL), server, liteDB
}

func afterTest(server *httptest.Server) {
	server.Close()
}

func TestToDoService_HomePageShowInfo(t *testing.T) {
	client, server, _ := beforeTest(t)

	client.GET("/").Expect().Status(200).JSON().Object().ContainsMap(map[string]string{
		"app":  "togo promax",
		"from": "Seth Phat",
	})

	afterTest(server)
}

func TestToDoService_LoginFailedBecauseEmptyRequest(t *testing.T) {
	client, server, _ := beforeTest(t)

	// 401
	client.POST("/login").Expect().Status(http.StatusUnauthorized)

	afterTest(server)
}

func TestToDoService_LoginSuccess(t *testing.T) {
	client, server, lite := beforeTest(t)

	email := faker.Email()
	userPassword := faker.Password()
	password, _ := helpers.HashPassword(userPassword)
	lite.DB.Exec(`INSERT INTO users (id, password, max_todo) VALUES (?, ?, ?)`, email, password, 5)

	// 200 with token
	client.POST("/login").
		WithFormField("user_id", email).
		WithFormField("password", userPassword).
		Expect().Status(200).
		JSON().Object().ContainsKey("data")

	afterTest(server)
}

func TestToDoService_AccessSecuredPathNeedAuthentication(t *testing.T) {
	client, server, _ := beforeTest(t)

	client.GET("/tasks").Expect().Status(401)

	afterTest(server)
}

func TestToDoService_ListTasks(t *testing.T) {
	client, server, lite := beforeTest(t)

	email := faker.Email()
	userPassword := faker.Password()
	password, _ := helpers.HashPassword(userPassword)
	lite.DB.Exec(`INSERT INTO users (id, password, max_todo) VALUES (?, ?, ?)`, email, password, 5)

	// 200 with token
	result := client.POST("/login").
		WithFormField("user_id", email).
		WithFormField("password", userPassword).
		Expect()

	result.Status(200).
		JSON().Object().ContainsKey("data")

	jwtToken := result.JSON().Object().Value("data").String().Raw()

	client.GET("/tasks").WithHeader("Authorization", jwtToken).
		Expect().Status(200).
		JSON().Object().ContainsKey("data")

	// add tasks
	lite.AddTask(
		context.Background(),
		&storages.Task{
			ID:          faker.UUIDHyphenated(),
			Content:     faker.Name(),
			UserID:      email,
			CreatedDate: "2020-05-05",
		},
	)

	client.GET("/tasks").WithHeader("Authorization", jwtToken).
		Expect().Status(200).
		JSON().Object().
		ContainsKey("data").Path("$.data").
		Array().First().Object().ContainsMap(map[string]string{
		"user_id":      email,
		"created_date": "2020-05-05",
	})

	afterTest(server)
}

func TestToDoService_AddTasks(t *testing.T) {
	client, server, lite := beforeTest(t)

	email := faker.Email()
	userPassword := faker.Password()
	password, _ := helpers.HashPassword(userPassword)
	lite.DB.Exec(`INSERT INTO users (id, password, max_todo) VALUES (?, ?, ?)`, email, password, 1)

	// 200 with token
	result := client.POST("/login").
		WithFormField("user_id", email).
		WithFormField("password", userPassword).
		Expect()

	result.Status(200).
		JSON().Object().ContainsKey("data")

	jwtToken := result.JSON().Object().Value("data").String().Raw()

	client.POST("/tasks").WithHeader("Authorization", jwtToken).
		WithJSON(map[string]string{
			"Content": "my new tasks",
		}).
		Expect().Status(200).
		JSON().Object().
		ContainsKey("data").Path("$.data").
		Object().ContainsMap(map[string]string{
		"user_id":      email,
		"created_date": time.Now().Format("2006-01-02"),
	})

	// can't add because reached maximum
	client.POST("/tasks").WithHeader("Authorization", jwtToken).
		WithJSON(map[string]string{
			"Content": "my new tasks 2",
		}).
		Expect().Status(500).
		JSON().Object().
		ContainsKey("error")

	afterTest(server)
}
