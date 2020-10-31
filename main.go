package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/manabie-com/togo/config"
	"github.com/manabie-com/togo/internal/storages/sqlite"
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

	db, err := sql.Open("sqlite3", cfg.TodoStore.LDB.Path)
	if err != nil {
		logging.Logger.Fatalw("error opening db", "detail", err)
	}
	newLiteDB := &sqlite.LiteDB{DB: db}
	todoUs := usecase.NewTogoUsecase(newLiteDB)

	mux := mux.InitWithLogger(logging.Logger.Desugar())

	transport.NewTogoHandler(mux, &todoUs, cfg.JWTKey)
	logging.Logger.Infof("Listening at %s", cfg.Address)
	http.ListenAndServe(":5050", mux)
}
