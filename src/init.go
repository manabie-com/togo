package server

import (
	"strings"

	"github.com/HoangMV/togo/lib/log"
	"github.com/HoangMV/togo/lib/pgsql"
	"github.com/HoangMV/togo/src/route"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
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
	app.Use(recover.New(recover.Config{EnableStackTrace: true}))
	app.Use(logger.New(logger.Config{Next: func(c *fiber.Ctx) bool {
		return strings.Contains(c.Request().URI().String(), "healthcheck")
	}}))

	r := route.New()
	r.Install(app)

	return app.Listen(sv.conf.Port)
}
