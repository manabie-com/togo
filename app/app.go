package app

import (
	"context"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/huuthuan-nguyen/manabie/app/handler"
	"github.com/huuthuan-nguyen/manabie/app/router"
	"github.com/huuthuan-nguyen/manabie/config"
	"github.com/huuthuan-nguyen/manabie/migration"
	"log"
	"net/http"
	"time"
)

type App struct {
	router  *mux.Router
	handler *handler.Handler
	ctx     context.Context
	server  *http.Server
	config  *config.Config
}

func (app *App) Run() {

	if err := app.Migrate(); err != nil {
		log.Fatalf("Migrating fail:%s\n", err)
		return
	}

	app.server = &http.Server{
		Addr:         fmt.Sprintf("%s:%s", app.config.Server.Host, app.config.Server.Port),
		Handler:      app.router,
		WriteTimeout: 30 * time.Second,
		ReadTimeout:  30 * time.Second,
	}
	defer app.server.Close()

	log.Printf("Listening on %s:%v...\n", app.config.Server.Host, app.config.Server.Port)
	log.Fatalln(app.server.ListenAndServe())
}

func NewApp(config *config.Config) *App {
	c := context.Background()
	h := handler.New(c, config)
	r := router.NewRouter(config, h)

	return &App{
		router:  r,
		handler: h,
		ctx:     c,
		config:  config,
	}
}

// Migrate /**
func (app *App) Migrate() error {
	migrateEngine := migration.NewEngine(app.config)
	return migrateEngine.Migrate()
}
