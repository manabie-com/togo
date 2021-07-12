package api

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/manabie-com/togo/config"
	"github.com/manabie-com/togo/internal/api/middleware"
	taskStorages "github.com/manabie-com/togo/internal/api/task/storages"
	taskPostgres "github.com/manabie-com/togo/internal/api/task/storages/postgres"
	taskSqlite "github.com/manabie-com/togo/internal/api/task/storages/sqlite"
	TaskTransport "github.com/manabie-com/togo/internal/api/task/transport"
	taskUseCase "github.com/manabie-com/togo/internal/api/task/usecase"
	userStorages "github.com/manabie-com/togo/internal/api/user/storages"
	userPostgres "github.com/manabie-com/togo/internal/api/user/storages/postgres"
	userSqlite "github.com/manabie-com/togo/internal/api/user/storages/sqlite"
	UserTransport "github.com/manabie-com/togo/internal/api/user/transport"
	userUseCase "github.com/manabie-com/togo/internal/api/user/usecase"
	"github.com/manabie-com/togo/internal/api/utils"
	"github.com/manabie-com/togo/internal/pkg/token/jwt"
	"net/http"
)

// ToDoHandler implement HTTP server
type ToDoHandler struct {
	UserTp *UserTransport.User
	TaskTp *TaskTransport.Task
}

// CreateAPIEngine creates engine instance that serves API endpoints,
func CreateAPIEngine(cfg *config.Config) (*http.Server, error) {
	userTp, taskTp, err := CreateTransport(cfg)
	if err != nil {
		return nil, err
	}

	generator := createJWTGenerator(cfg)
	userTp.TokenGenerator = generator

	handler := ToDoHandler{
		UserTp: userTp,
		TaskTp: taskTp,
	}

	apiDomainString := fmt.Sprintf("%v:%v", cfg.RestfulAPI.Host, cfg.RestfulAPI.Port)
	server := &http.Server{Addr: apiDomainString, Handler: middleware.AddCors(middleware.ValidToken(&handler, cfg, generator))}
	return server, nil
}

func initSQLiteDB(cfg *config.Config) (*sql.DB, error) {
	return sql.Open("sqlite3", cfg.DBs.SQLite.DataSourceName)
}

func initPostgresDB(cfg *config.Config) (*sql.DB, error) {
	connectionString := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		cfg.DBs.Postgres.Host, cfg.DBs.Postgres.Port, cfg.DBs.Postgres.Username, cfg.DBs.Postgres.Password, cfg.DBs.Postgres.Database)
	return sql.Open("postgres", connectionString)
}

func CreateTransport(cfg *config.Config) (*UserTransport.User, *TaskTransport.Task, error) {

	var userStore userStorages.Store
	var taskStore taskStorages.Store
	var db *sql.DB
	var err error

	switch cfg.Store {
	case "sqlite":
		db, err = initSQLiteDB(cfg)
		userStore = &userSqlite.LiteDB{DB: db}
		taskStore = &taskSqlite.LiteDB{DB: db}
	case "postgres":
		db, err = initPostgresDB(cfg)
		userStore = &userPostgres.PostgresDB{DB: db}
		taskStore = &taskPostgres.PostgresDB{DB: db}
	default:
		err = errors.New("empty db type")
	}
	if err != nil {
		return nil, nil, err
	}

	userUC := userUseCase.User{
		Store: userStore,
	}
	taskUC := taskUseCase.Task{
		Store:           taskStore,
		UserStore:       userStore,
		GeneratorUUIDFn: utils.GenerateNewUUID,
	}

	return &UserTransport.User{
			UserUC: userUC,
		}, &TaskTransport.Task{
			TaskUC: taskUC,
		}, nil
}

func createJWTGenerator(cfg *config.Config) *jwt.Generator {
	return &jwt.Generator{
		Cfg: cfg,
	}
}

func (s *ToDoHandler) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	switch req.URL.Path {
	case "/login":
		switch req.Method {
		case http.MethodPost:
			s.UserTp.Login(resp, req)
			return
		}
	case "/tasks":
		switch req.Method {
		case http.MethodGet:
			s.TaskTp.List(resp, req)
			return
		case http.MethodPost:
			s.TaskTp.Add(resp, req)
			return
		}
	default:
		http.NotFound(resp, req)
		return
	}
	http.Error(resp, "405 method not allowed", http.StatusMethodNotAllowed)
	return
}
