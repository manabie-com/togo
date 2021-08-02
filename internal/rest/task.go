package rest

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"time"
	"togo/config"
	"togo/internal/dto"
	"togo/internal/repository"
	"togo/internal/service"
)

//type SearchTasksRequest struct {
//	IsDone      *bool      `query:"is_done"`
//	CreatedDate *time.Time `query:"created_date"`
//}

func ListTask(sc *config.ServerConfig) fiber.Handler {
	return func(c *fiber.Ctx) error {
		currentUser := GetCurrentUser(c)

		//var r SearchTasksRequest
		//
		//if err := c.QueryParser(&r); err != nil {
		//	fmt.Println(err)
		//	return SimpleError(c, err)
		//}

		db := sc.DB
		repo := repository.NewRepo(db)
		svc := service.NewTaskService(repo)

		tasks, err := svc.ListTasks(c.UserContext(), currentUser.ID)
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

		db := sc.DB
		repo := repository.NewRepo(db)
		svc := service.NewTaskService(repo)

		task, err := svc.Create(c.UserContext(), currentUser, data.Content, currentUser.ID, time.Now())
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
		db := sc.DB
		repo := repository.NewRepo(db)
		svc := service.NewTaskService(repo)

		task, err := svc.GetTask(c.UserContext(), int32(id), currentUser.ID)
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
		db := sc.DB
		repo := repository.NewRepo(db)
		svc := service.NewTaskService(repo)

		err = svc.DeleteTask(c.UserContext(), int32(id), currentUser.ID)
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
