package postgres_test

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"
	"testing"
	"time"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/manabie-com/togo/internal/entities"
	"github.com/manabie-com/togo/internal/storages/postgres"
	"github.com/ory/dockertest/v3"
	"github.com/stretchr/testify/assert"
)

var (
	pdb postgres.PDB
	now = time.Now().Format("2006-01-02")
)

func TestRetrieveTasks(t *testing.T) {
	var tests = []struct {
		userID         string
		createdDate    string
		expectedStatus bool // isGotTaskSuccess
	}{
		{
			"firstUser",
			"2020-08-11",
			true,
		},
		{

			"firstUser",
			"2020-08-22",
			false,
		},
		{
			"secUser",
			"2020-08-22",
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.userID, func(t *testing.T) {
			tasks, err := pdb.RetrieveTasks(context.Background(), convertNullString(tt.userID), convertNullString(tt.createdDate))
			assert.NoError(t, err)
			assert.Equal(t, tt.expectedStatus, tasks != nil)
		})
	}
}
func TestValidateUser(t *testing.T) {
	var tests = []struct {
		userID        string
		password      string
		expectedValid bool
	}{
		{
			"firstUser",
			"example",
			true,
		},
		{
			"firstUser",
			"examplee",
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.userID, func(t *testing.T) {
			isValid := pdb.ValidateUser(context.Background(), convertNullString(tt.userID), convertNullString(tt.password))
			assert.Equal(t, tt.expectedValid, isValid)
		})
	}
}

func TestGetHashedPass(t *testing.T) {
	var tests = []struct {
		userID   string
		expected bool //isGotPasswordSuccess and no error
	}{
		{
			"firstUser",
			true,
		},
		{
			"firstUser1",
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.userID, func(t *testing.T) {
			hashedPass, err := pdb.GetHashedPass(context.Background(), convertNullString(tt.userID))
			assert.Equal(t, tt.expected, err == nil) // check if doesn't have error or not
			assert.Equal(t, tt.expected, hashedPass != "")
		})
	}
}

func TestGetMaxTaskTodo(t *testing.T) {
	var tests = []struct {
		userID   string
		expected bool // isGetSuccess
	}{
		{
			"firstUser",
			true,
		},
		{
			"firstUser1",
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.userID, func(t *testing.T) {
			maxTask, err := pdb.GetMaxTaskTodo(context.Background(), tt.userID)
			assert.Equal(t, tt.expected, err == nil) // check if doesn't have error or not
			assert.Equal(t, tt.expected, maxTask != 0)
		})
	}
}

func TestAddTask(t *testing.T) {
	var tests = []struct {
		task     entities.Task
		expected bool // isSucceed
	}{
		{
			entities.Task{ID: "e35e13f8-35f3-409f-8e2f-f3e0373f1ca3", Content: "content", CreatedDate: now, UserID: "firstUser"},
			true,
		},
		{
			entities.Task{ID: "e35e13f8-35f3-409f-8e2f-f3e0373f1ca3", Content: "content", CreatedDate: now, UserID: "firstUser1"},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.task.UserID, func(t *testing.T) {
			err := pdb.AddTask(context.Background(), tt.task)
			assert.Equal(t, tt.expected, assert.NoError(t, err))
		})
	}
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
			"file://../../../migrations/postgres", // depend on your migrations
			fmt.Sprintf("postgres://postgres:secret@localhost:%s/%s?sslmode=disable", resource.GetPort("5432/tcp"), dbName),
		)
		pdb = postgres.PDB{DB: Pool}
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
func convertNullString(s string) sql.NullString {
	if s == "" {
		return sql.NullString{}
	}
	return sql.NullString{String: s, Valid: true}
}
