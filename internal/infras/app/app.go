package app

import (
	"fmt"

	"github.com/datshiro/togo-manabie/internal/infras/db"
	"github.com/datshiro/togo-manabie/internal/interfaces"
	"github.com/labstack/echo/v4"
)

type App interface {
	ConfigMiddleware()
	ConfigLogLevel()
	ConfigErrHandler()
	ConfigLogFormat()
	RegisterHandlers()
	Parse()
	Run() error
}

func NewApp() App {
	return &app{e: echo.New()}
}

type app struct {
	e             *echo.Echo
	LogLevel      string
	DbUrl         string
	Host          string
	ClientHost    string
	ApiPrefix     string
	Port          string
	RedisLocation string
}

func (a *app) RegisterHandlers() {
	dbc, err := db.NewPGDB(a.DbUrl)
	if err != nil {
		fmt.Println(err)
		panic("Invalid database connection")
	}

	interfaces.RegisterHandlers(a.e, a.ApiPrefix, dbc)
}

func (a *app) Run() error {
	address := a.Host + ":" + a.Port

	return a.e.Start(address)
}
