package transport

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"time"
	"togo/config"
	"togo/internal/domain"
	"togo/internal/dto"
	"togo/internal/redix"
	"togo/internal/repository"
)

func ListTask(sc *config.ServerConfig) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var r dto.SearchTasksRequest

		if err := c.QueryParser(&r); err != nil {
			return SimpleError(c, err)
		}

		currentUser := GetCurrentUser(c)

		rdb := sc.Redis
		db := sc.DB
		repo := repository.NewRepo(db)
		rdbStore := redix.NewRedisStore(rdb)
		taskDomain := domain.NewTaskDomain(repo, rdbStore)

		tasks, err := taskDomain.ListTasks(c.UserContext(), currentUser.ID, r.IsDone, r.CreatedDate)
		if err != nil {
			return SimpleError(c, err)
		}

		return c.Status(200).JSON(tasks)
	}
}

func CreateTask(sc *config.ServerConfig) fiber.Handler {
	return func(c *fiber.Ctx) error {
		data := new(dto.CreateTaskDTO)

		if err := c.BodyParser(data); err != nil {
			return SimpleError(c, err)
		}

		if err := validator.New().Struct(data); err != nil {
			return SimpleError(c, err)
		}

		currentUser := GetCurrentUser(c)

		rdb := sc.Redis
		db := sc.DB
		repo := repository.NewRepo(db)
		rdbStore := redix.NewRedisStore(rdb)
		taskDomain := domain.NewTaskDomain(repo, rdbStore)

		task, err := taskDomain.Create(c.UserContext(), *currentUser, data.Content, currentUser.ID, time.Now())
		if err != nil {
			return SimpleError(c, err)
		}

		return c.Status(200).JSON(task)
	}
}

func GetTask(sc *config.ServerConfig) fiber.Handler {
	return func(c *fiber.Ctx) error {
		currentUser := GetCurrentUser(c)

		id, err := strconv.ParseInt(c.Params("id"), 10, 32)
		if err != nil {
			return SimpleError(c, err)
		}

		rdb := sc.Redis
		db := sc.DB
		repo := repository.NewRepo(db)
		rdbStore := redix.NewRedisStore(rdb)
		taskDomain := domain.NewTaskDomain(repo, rdbStore)

		task, err := taskDomain.GetTask(c.UserContext(), int32(id), currentUser.ID)
		if err != nil {
			return SimpleError(c, err)
		}

		return c.Status(200).JSON(task)
	}
}

func DeleteTask(sc *config.ServerConfig) fiber.Handler {
	return func(c *fiber.Ctx) error {
		currentUser := GetCurrentUser(c)

		id, err := strconv.ParseInt(c.Params("id"), 10, 32)
		if err != nil {
			return SimpleError(c, err)
		}

		rdb := sc.Redis
		db := sc.DB
		repo := repository.NewRepo(db)
		rdbStore := redix.NewRedisStore(rdb)
		taskDomain := domain.NewTaskDomain(repo, rdbStore)

		err = taskDomain.DeleteTask(c.UserContext(), int32(id), *currentUser)
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
		data := new(dto.UpdateTaskDTO)

		if err := c.BodyParser(data); err != nil {
			return SimpleError(c, err)
		}

		v := validator.New()
		err := v.Struct(data)
		if err != nil {
			return SimpleError(c, err)
		}

		currentUser := GetCurrentUser(c)

		id, err := strconv.ParseInt(c.Params("id"), 10, 32)
		if err != nil {
			return SimpleError(c, err)
		}

		rdb := sc.Redis
		db := sc.DB
		repo := repository.NewRepo(db)
		rdbStore := redix.NewRedisStore(rdb)
		taskDomain := domain.NewTaskDomain(repo, rdbStore)

		err = taskDomain.UpdateTask(c.UserContext(), *currentUser, int32(id), data.IsDone)
		if err != nil {
			return SimpleError(c, err)
		}

		return c.Status(200).JSON(&fiber.Map{
			"msg": "OK",
		})
	}
}
