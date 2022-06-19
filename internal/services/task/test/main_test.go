package test

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/go-pg/pg/v10"
	"github.com/sirupsen/logrus"
	"testing"
	"togo/internal/pkg/db"
	"togo/internal/services/task/api"
	"togo/internal/services/task/application"
	"togo/internal/services/task/store/postgres"
)

var userRepo *postgres.UserRepository
var taskRepo *postgres.TaskRepository

const connectionString = "postgres://postgres:postgres@localhost:5432/togo?sslmode=disable"

var server *gin.Engine

func TestMain(m *testing.M) {
	server = gin.New()
	server.Use(gin.Recovery())
	postgresDB := &db.DB{}
	opt, err := pg.ParseURL(connectionString)
	if err != nil {
		logrus.Panic("parse go-pg connection string err: %w", err)
	}

	postgresDB.Conn = pg.Connect(opt)
	ctx := context.Background()

	if err = postgresDB.Conn.Ping(ctx); err != nil {
		panic(err)
	}
	userRepo = postgres.NewUserRepository(postgresDB)
	taskRepo = postgres.NewTaskRepository(postgresDB)

	userService := application.NewUserService(userRepo)
	taskService := application.NewTaskService(taskRepo, userRepo)

	userApi := api.NewUserAPI(userService)
	taskApi := api.NewTaskAPI(taskService)

	userApi.AddRoutes(server)
	taskApi.AddRoutes(server)

	m.Run()
}
