package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/manabie-com/togo/config"
	"github.com/manabie-com/togo/internal/storages/postgres"
	"github.com/manabie-com/togo/internal/transport"
	"github.com/manabie-com/togo/internal/usecase"
	"github.com/manabie-com/togo/logging"

	"github.com/manabie-com/togo/mux"
	_ "github.com/mattn/go-sqlite3"
	"go.uber.org/zap"
)

func main() {
	var err error
	// load the config, load the file default, and load the environment variable
	cfg := config.LoadConfigFile("config.yml")
	config.LoadConfigEnv(&cfg)

	// logging for togo
	logging.Logger, err = logging.InitWithOption(cfg.LogLevel, cfg.Service)
	if err != nil {
		log.Println("can't setup zap log", err)
	}
	zap.ReplaceGlobals(logging.Logger.Desugar())
	defer logging.Logger.Sync()

	// connect DB
	pool, err := NewPDBConnection(context.Background(), dsnPDBGenerator(cfg))
	if err != nil {
		logging.Logger.Panicw("error opening db", "detail", err)
	}
	storages := postgres.PDB{DB: pool}

	todoUs := usecase.NewTogoUsecase(storages)

	mux := mux.InitWithLogger(logging.Logger.Desugar())
	transport.NewTogoHandler(mux, todoUs, cfg.JWTKey)
	logging.Logger.Infof("Listening at %s", cfg.Address)
	http.ListenAndServe(":5050", mux)
}

func dsnPDBGenerator(cfg config.Config) string {
	return fmt.Sprintf("postgresql://%s:%s@%s:%s/%s", cfg.TodoStore.PDB.Username, cfg.TodoStore.PDB.Password, cfg.TodoStore.PDB.Host, cfg.TodoStore.PDB.Port, cfg.TodoStore.PDB.DbName)
}

//NewPDBConnection create pool connection for postgres
func NewPDBConnection(ctx context.Context, dns string) (*pgxpool.Pool, error) {
	return pgxpool.Connect(ctx, dns)
}

//NewSQLiteConnection create new connection for SQLite
func NewSQLiteConnection(path string) (*sql.DB, error) {
	return sql.Open("sqlite3", path)

}
