package main

import (
	"context"
	"time"
	"togo/configs"
	"togo/internal/provider"
	"togo/internal/provider/hasher"
	"togo/internal/provider/jwt"
	"togo/internal/repository"
	"togo/internal/repository/gormrepo"
	"togo/internal/repository/redisrepo"
	"togo/internal/services"
	"togo/internal/transport"
	"togo/internal/transport/http"
	"togo/internal/usecase"

	"github.com/go-redis/redis/v8"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	config *configs.Config

	database *gorm.DB
	redisDB  *redis.Client

	bcryptProvider provider.PasswordHashProvider
	jwtProvider    provider.TokenProvider

	userRepo      repository.UserRepository
	taskRepo      repository.TaskRepository
	taskLimitRepo repository.TaskLimitRepository

	userUsecase usecase.UserUsecase
	authUsecase usecase.AuthUsecase
	taskUsecase usecase.TaskUsecase

	server transport.Server
)

func initConfig() {
	config = configs.NewConfig()
}

func initProviders() {
	bcryptProvider = hasher.NewBcryptProvider(config.EncryptSalt, 10)
	jwtProvider = jwt.NewJWTProvider(config.JwtSigningKey, config.JwtAccessTokenDuration, config.JwtRefreshTokenDuration)
}

func initStorages(ctx context.Context) (err error) {
	redisDB = redis.NewClient(&redis.Options{
		Addr:     config.RedisAddr,
		// Password: config.RedisPassword,
		// DB:       config.RedisDB,
	})

	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()
	database, err = gorm.Open(postgres.Open(config.DatabaseURI))
	return
}

func initRepositories() {
	userRepo = gormrepo.NewUserRepository(database)
	taskRepo = gormrepo.NewTaskRepository(database)
	taskLimitRepo = redisrepo.NewTaskLimitRepository(redisDB, "task_limit")
}

func initServices() {
	userUsecase = services.NewUserService(bcryptProvider, userRepo)
	authUsecase = services.NewAuthService(bcryptProvider, jwtProvider, userRepo)
	taskUsecase = services.NewTaskService(userRepo, taskRepo, taskLimitRepo)
}

func initTransport(ctx context.Context) error {
	server = http.NewHTTPServer(userUsecase, authUsecase, taskUsecase)
	if err := server.Load(ctx); err != nil {
		return err
	}
	return nil
}

func loadApp(ctx context.Context) error {
	initConfig()
	initProviders()
	if err := initStorages(ctx); err != nil {
		return err
	}
	initRepositories()
	initServices()
	if err := initTransport(ctx); err != nil {
		return err
	}
	return nil
}

func start() {
	server.Serve(config.Host, config.Port)
}

func main() {
	ctx := context.Background()
	if err := loadApp(ctx); err != nil {
		panic(err)
	}
	start()
}
