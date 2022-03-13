package route

import (
	"github.com/HoangMV/togo/src/biz"
	"github.com/gofiber/fiber/v2"
)

type Route struct {
	biz *biz.Biz
}

func New() *Route {
	return &Route{biz.New()}
}

func (r *Route) Install(app *fiber.App) {

	v1 := app.Group("/api/v1")
	v1.Get("/healthcheck/status")

	v1.Post("/register")
	v1.Post("/login")

	v1.Post("/todo")
}
