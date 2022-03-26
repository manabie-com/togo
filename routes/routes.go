package routes

import (
	"github.com/luongdn/togo/models"
	"github.com/luongdn/togo/storage"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func SetUpRoutes(app *fiber.App) {
	app.Use(logger.New())
	app.Get("/", Hello)

	app.Post("/users", CreateUser)
	app.Get("/users/:user_id/tasks", ListTasks)
	app.Post("/users/:user_id/tasks", RateLimiter(models.TaskCreate), CreateTask)

	app.Post("/users/:user_id/rules", CreateRule)
	app.Put("/users/:user_id/rules", UpdateRule)

	app.Use(cors.New())

	app.Use(func(ctx *fiber.Ctx) error {
		return ctx.SendStatus(404)
	})
}

func Hello(ctx *fiber.Ctx) error {
	return ctx.SendString("Hello, World!")
}

func CreateUser(ctx *fiber.Ctx) error {
	user := new(models.User)
	if err := ctx.BodyParser(user); err != nil {
		return ctx.Status(400).JSON(err.Error())
	}

	store := storage.GetUserStore()
	err := store.CreateUser(ctx.Context(), user)
	if err != nil {
		return ctx.Status(400).JSON(err.Error())
	}
	return ctx.Status(200).JSON(user)
}

func ListTasks(ctx *fiber.Ctx) error {
	store := storage.GetTaskStore()
	tasks, err := store.ListTasks(ctx.Context(), ctx.Params("user_id"))
	if err != nil {
		return ctx.Status(400).JSON(err.Error())
	}
	return ctx.Status(200).JSON(tasks)
}

func CreateTask(ctx *fiber.Ctx) error {
	task := new(models.Task)
	if err := ctx.BodyParser(task); err != nil {
		return ctx.Status(400).JSON(err.Error())
	}

	store := storage.GetTaskStore()
	err := store.CreateTask(ctx.Context(), ctx.Params("user_id"), task)
	if err != nil {
		return ctx.Status(400).JSON(err.Error())
	}
	return ctx.Status(200).JSON(task)
}

func CreateRule(ctx *fiber.Ctx) error {
	rule := new(models.Rule)
	if err := ctx.BodyParser(rule); err != nil {
		return ctx.Status(400).JSON(err.Error())
	}

	store := storage.GetRuleStore()
	err := store.CreateRule(ctx.Context(), ctx.Params("user_id"), rule)
	if err != nil {
		return ctx.Status(400).JSON(err.Error())
	}
	return ctx.Status(200).JSON(rule)
}

func UpdateRule(ctx *fiber.Ctx) error {
	rule := new(models.Rule)
	if err := ctx.BodyParser(rule); err != nil {
		return ctx.Status(400).JSON(err.Error())
	}

	store := storage.GetRuleStore()
	err := store.UpdateRule(ctx.Context(), ctx.Params("user_id"), rule)
	if err != nil {
		return ctx.Status(400).JSON(err.Error())
	}
	return ctx.Status(200).JSON(rule)
}
