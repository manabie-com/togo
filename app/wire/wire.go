//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.
package wire

import (
	"ansidev.xyz/pkg/db"
	"ansidev.xyz/pkg/rds"
	"database/sql"
	authRepository "github.com/ansidev/togo/auth/repository"
	authService "github.com/ansidev/togo/auth/service"
	taskRepository "github.com/ansidev/togo/task/repository"
	taskService "github.com/ansidev/togo/task/service"
	"github.com/go-redis/redis/v8"
	"time"

	userRepository "github.com/ansidev/togo/user/repository"
	"github.com/google/wire"
	"gorm.io/gorm"
)

func InitSqlClient(config db.SqlDbConfig) *sql.DB {
	wire.Build(db.NewPostgresClient)

	return &sql.DB{}
}

func InitRedisClient(config rds.RedisConfig) *redis.Client {
	wire.Build(rds.NewRedisClient)

	return &redis.Client{}
}

func InitAuthService(db *gorm.DB, rdb *rds.RedisDB, tokenTTL time.Duration) authService.IAuthService {
	wire.Build(authService.NewAuthService, userRepository.NewPostgresUserRepository, authRepository.NewRedisCredentialRepository)
	return &authService.AuthService{}
}

func InitTaskService(db *gorm.DB) taskService.ITaskService {
	wire.Build(taskService.NewTaskService, userRepository.NewPostgresUserRepository, taskRepository.NewPostgresTaskRepository)
	return &taskService.TaskService{}
}
