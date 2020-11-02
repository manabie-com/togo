package togo_test

import (
	"bytes"
	"context"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/go-chi/chi"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/manabie-com/togo/internal/entities"
	"github.com/manabie-com/togo/internal/storages/postgres"
	"github.com/manabie-com/togo/internal/transport"
	"github.com/manabie-com/togo/internal/usecase"
	"github.com/ory/dockertest/v3"
	"github.com/stretchr/testify/assert"
)

var (
	Storages postgres.PDB
	mux      *chi.Mux
)

const jwt_key = "wqGyEBBfPK9w3Lxw"

func TestLogin(t *testing.T) {
	var tests = []struct {
		url            string
		jsonRequest    []byte
		expectedStatus int
	}{
		{
			"/login",
			[]byte(`{"user_id": "firstUser","password": "example"}`),
			200,
		},
		{
			"/login/",
			[]byte(`{"user_id": "firstUser","password": "example"}`),
			200,
		},
		{
			"/login//",
			[]byte(`{"user_id": "firstUser","password": "example"}`),
			404,
		},
		{
			"/login?notexsitusername",
			[]byte(`{"user_id": "manabie","password": "example"}`),
			401,
		},
		{
			"/login?notexsitusername",
			[]byte(`{"user_id": "firstUser","password": "example"}`),
			200,
		},
		{
			"/login?",
			[]byte(`{"user_id": "firstUser","password": "example"}`),
			200,
		},
		{
			"/login",
			[]byte(`{"user_id": "firstUser","password": "wrong pass"}`),
			401,
		},
		{
			"/loginn",
			[]byte(`{"user_id": "firstUser","password": "example"}`),
			404,
		},
		{
			"/login.",
			[]byte(`{"user_id": "firstUser","password": "example"}`),
			404,
		},
	}

	for _, tt := range tests {
		t.Run(tt.url, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodPost, tt.url, bytes.NewBuffer(tt.jsonRequest))
			rec := httptest.NewRecorder()
			mux.ServeHTTP(rec, req)
			assert.Equal(t, tt.expectedStatus, rec.Code)
		})
	}
}

func TestListContent(t *testing.T) {
	var tests = []struct {
		url                 string
		authorizationHeader string
		expectedStatus      int
	}{
		{
			fmt.Sprintf("/tasks?created_date=2020-11-01"),
			createToken("firstUser", 15),
			200,
		},
		{
			fmt.Sprintf("/tasks?created_date=2020-11-01"),
			createToken("firstUser", -1), //invalid token, token is expired
			401,
		},
		{
			fmt.Sprintf("/tasks?created_date=2020-11-01?user"),
			createToken("firstUser", 15), //invalid token, token is expired,
			400,
		},
		{
			fmt.Sprintf("/tasks?created_date=2020-11-01/"),
			createToken("firstUser", 15), //invalid token, token is expired,
			400,
		},
		{
			fmt.Sprintf("/tasks?created_date=2020-11-01//"),
			createToken("firstUser", 15), //invalid token, token is expired,
			400,
		},
		{
			fmt.Sprintf("/tasks"),
			createToken("firstUser", 15), //invalid token, token is expired
			400,
		},
	}
	for _, tt := range tests {
		t.Run(tt.url, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodGet, tt.url, nil)
			req.Header.Set("Authorization", tt.authorizationHeader)
			rec := httptest.NewRecorder()
			mux.ServeHTTP(rec, req)
			assert.Equal(t, tt.expectedStatus, rec.Code)
		})
	}
}
func TestAddTask(t *testing.T) {
	var tests = []struct {
		url                 string
		jsonRequest         []byte
		authorizationHeader string
		expectedStatus      int
	}{
		{
			"/tasks",
			[]byte(`{"content": "this is task content"}`),
			createToken("firstUser", 15),
			200,
		},
		{
			"/tasks?",
			[]byte(`{"content": "this is task content"}`),
			createToken("firstUser", 15),
			200,
		},
		{
			"/tasks/",
			[]byte(`{"content": "this is task content"}`),
			createToken("firstUser", 15),
			200,
		},
		{
			"/tasks//",
			[]byte(`{"content": "this is task content"}`),
			createToken("firstUser", 15),
			404,
		},
		{
			"/tasks",
			[]byte(`{"contentt": "this is task content"}`),
			createToken("firstUser", 15),
			400,
		},
		//test rate limit
		{
			"/tasks",
			[]byte(`{"content": "this is task content"}`),
			createToken("firstUser", 15),
			403,
		},
	}
	for idx, tt := range tests {
		t.Run(tt.url, func(t *testing.T) {
			// the last test to test rate limiting (limit n task per user per day)
			if idx == len(tests)-1 {
				// add max_todo task to make sure exceeding the rate litmit
				maxTodo, err := Storages.GetMaxTaskTodo(context.Background(), "firstUser")
				if err != nil {
					t.Fatal("Can't get number of task", err)
				}
				for i := 0; i < maxTodo; i++ {
					req := httptest.NewRequest(http.MethodPost, tests[idx].url, bytes.NewBuffer([]byte(tests[idx].jsonRequest)))
					req.Header.Set("Authorization", tests[idx].authorizationHeader)
					rec := httptest.NewRecorder()
					mux.ServeHTTP(rec, req)
				}
			}
			req := httptest.NewRequest(http.MethodPost, tt.url, bytes.NewBuffer([]byte(tt.jsonRequest)))
			req.Header.Set("Authorization", tt.authorizationHeader)
			rec := httptest.NewRecorder()
			mux.ServeHTTP(rec, req)
			assert.Equal(t, tt.expectedStatus, rec.Code)
		})
	}
}
func addTask(userID, content string) error {
	task := entities.Task{}
	now := time.Now()
	task.Content = content
	task.ID = uuid.New().String()
	task.UserID = userID
	task.CreatedDate = now.Format("2006-01-02")
	err := Storages.AddTask(context.Background(), task)
	return err
}
func createToken(userID string, duration int) string {
	atClaims := jwt.MapClaims{}
	atClaims["user_id"] = userID
	atClaims["exp"] = time.Now().Add(time.Minute * time.Duration(duration)).Unix()
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	token, err := at.SignedString([]byte(jwt_key))
	if err != nil {
		return ""
	}
	return token
}
func TestMain(m *testing.M) {
	pool, err := dockertest.NewPool("")
	if err != nil {
		log.Printf("Could not start new pool: %s", err)
	}
	dbName := "togo"
	resource, err := pool.Run("postgres", "13", []string{"POSTGRES_PASSWORD=secret", "POSTGRES_DB=" + dbName})

	if err != nil {
		log.Printf("Could not start resource: %s", err)
	}
	if err = pool.Retry(func() error {
		Pool, err := pgxpool.Connect(context.Background(), fmt.Sprintf("postgresql://postgres:secret@localhost:%s/%s?sslmode=disable", resource.GetPort("5432/tcp"), dbName))
		if err != nil {
			log.Printf("Could not connect resource: %s", err)
			return err
		}
		migrate, err := migrate.New(
			"file://../../migrations/postgres", // depend on your migrations
			fmt.Sprintf("postgres://postgres:secret@localhost:%s/%s?sslmode=disable", resource.GetPort("5432/tcp"), dbName),
		)
		mux = chi.NewRouter()
		Storages = postgres.PDB{DB: Pool}
		todoUs := usecase.NewTogoUsecase(Storages)
		transport.NewTogoHandler(mux, todoUs, jwt_key)
		if err != nil {
			log.Print("Could not migrate", err)
			return err
		}
		if err := migrate.Up(); err != nil {
			log.Println(err)
			return err
		}
		return nil
	}); err != nil {
		log.Printf("Could not connect to docker: %s", err)
	}
	if err != nil {
		log.Printf("Could not connect resource: %s", err)
	}
	code := m.Run()
	// When you're done, kill and remove the container
	if err = pool.Purge(resource); err != nil {
		log.Printf("Could not purge resource: %s", err)
	}
	os.Exit(code)
}
