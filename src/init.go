package server

import (
	"github.com/HoangMV/togo/lib/log"
	"github.com/HoangMV/togo/lib/pgsql"
	"github.com/HoangMV/togo/src/route"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

type ApiServer struct {
	conf *Config
}

func New() *ApiServer {
	conf := getConfigFromEnv()
	return &ApiServer{conf}
}

func (sv *ApiServer) SetupConfig() {
	log.Install()
	pgsql.Install()
}

func (sv *ApiServer) Run() error {
	app := fiber.New()
	app.Use(cors.New())

	r := route.New()
	r.Install(app)

	return app.Listen(sv.conf.Port)
}
