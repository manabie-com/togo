package route

import (
	"github.com/HoangMV/togo/lib/utils"
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
	v1.Get("/healthcheck/status", r.HealthCheck)

	v1.Post("/register", r.Register)
	v1.Post("/login", r.Login)

	v1.Post("/todo", r.CheckAuth, r.CreateUserTodo)
	v1.Put("/todo", r.CheckAuth, r.UpdateUserTodo)
	v1.Get("/todo", r.CheckAuth, r.GetUserListTodo)
}

func (r *Route) HealthCheck(ctx *fiber.Ctx) error {
	return utils.WriteSuccessEmptyContent(ctx)
}
