package main

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/spf13/cobra"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"log"
	"manabieAssignment/cmd/migrations"
	"manabieAssignment/config"
	todoRepository "manabieAssignment/internal/todo/repository"
	"manabieAssignment/internal/todo/transport"
	"manabieAssignment/internal/todo/usecase"
	userRepository "manabieAssignment/internal/user/repository"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	cfg, err := config.Load()
	if err != nil {
		log.Fatal(err)
	}

	cmd := &cobra.Command{
		Use: "server",
	}

	cmd.AddCommand(&cobra.Command{
		Use: "server",
		Run: func(cmd *cobra.Command, args []string) {
			err := serveServer(cfg)
			if err != nil {
				log.Fatal(err)
			}
		},
	}, migrations.MigrateCommand(cfg.DB.MigrationFolder, cfg.DB.Source),
	)
	return cmd.Execute()
}

func connectDB(cfg config.DBConfig) (*gorm.DB, error) {
	sqlDB, err := sql.Open(cfg.Driver, cfg.Source)
	if err != nil {
		return nil, err
	}
	gormDB, err := gorm.Open(postgres.New(postgres.Config{
		Conn: sqlDB,
	}), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
		NamingStrategy: schema.NamingStrategy{
			SingularTable: false,
		},
	})
	if err != nil {
		return nil, err
	}
	return gormDB, nil
}

func serveServer(cfg *config.Config) error {
	db, err := connectDB(cfg.DB)
	if err != nil {
		log.Fatalf("cannot connect to database %v", err)
	}
	r := gin.Default()

	todoRepo := todoRepository.NewTodoRepository(db)
	userRepo := userRepository.NewUserRepository(db)
	todoUC := usecase.NewTodoUseCase(todoRepo, userRepo)
	transport.NewTodoHandler(r, todoUC)
	return r.Run()
}
