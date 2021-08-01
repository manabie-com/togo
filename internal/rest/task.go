package rest

import (
	"github.com/gofiber/fiber/v2"
	"strconv"
	"time"
	"togo/config"
	"togo/internal/repository"
	"togo/internal/service"
)

func ListTask(sc *config.ServerConfig) fiber.Handler {
	return func(c *fiber.Ctx) error {
		userId := GetCurrentUserId(c)

		db := sc.DB
		repo := repository.NewRepo(db)
		svc := service.NewTaskService(repo)

		tasks, err := svc.ListTasks(c.UserContext(), userId)
		if err != nil {
			return SimpleError(c, err)
		}

		return c.Status(200).JSON(tasks)
	}
}

type CreateTaskDTO struct {
	Content string `json:"content" validate:"required,min=6,max=32"`
}

func CreateTask(sc *config.ServerConfig) fiber.Handler {
	return func(c *fiber.Ctx) error {
		userId := GetCurrentUserId(c)

		db := sc.DB
		repo := repository.NewRepo(db)
		svc := service.NewTaskService(repo)

		task, err := svc.Create(c.UserContext(), "ahih", userId, time.Now())
		if err != nil {
			return SimpleError(c, err)
		}

		return c.Status(200).JSON(task)
	}
}

func GetTask(sc *config.ServerConfig) fiber.Handler {
	return func(c *fiber.Ctx) error {
		userId := GetCurrentUserId(c)

		id, err := strconv.ParseInt(c.Params("id"), 10, 32)
		if err != nil {
			return SimpleError(c, err)
		}
		db := sc.DB
		repo := repository.NewRepo(db)
		svc := service.NewTaskService(repo)

		task, err := svc.GetTask(c.UserContext(), int32(id), userId)
		if err != nil {
			return SimpleError(c, err)
		}

		return c.Status(200).JSON(task)
	}
}

func DeleteTask(sc *config.ServerConfig) fiber.Handler {
	return func(c *fiber.Ctx) error {
		userId := GetCurrentUserId(c)

		id, err := strconv.ParseInt(c.Params("id"), 10, 32)
		if err != nil {
			return SimpleError(c, err)
		}
		db := sc.DB
		repo := repository.NewRepo(db)
		svc := service.NewTaskService(repo)

		err = svc.DeleteTask(c.UserContext(), int32(id), userId)
		if err != nil {
			return SimpleError(c, err)
		}

		return c.Status(200).JSON(&fiber.Map{
			"msg": "OK",
		})
	}
}
func UpdateTask(sc *config.ServerConfig) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id, err := strconv.ParseInt(c.Params("id"), 10, 32)
		if err != nil {
			return SimpleError(c, err)
		}
		db := sc.DB
		repo := repository.NewRepo(db)
		svc := service.NewTaskService(repo)

		err = svc.UpdateTask(c.UserContext(), int32(id), true)
		if err != nil {
			return SimpleError(c, err)
		}

		return c.Status(200).JSON(&fiber.Map{
			"msg": "OK",
		})
	}
}
