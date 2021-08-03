package transport

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"togo/config"
	"togo/internal/domain"
	"togo/internal/dto"
	"togo/internal/redix"
	"togo/internal/repository"
	"togo/utils"
)

func UserLogin(sc *config.ServerConfig) fiber.Handler {
	return func(c *fiber.Ctx) error {
		data := new(dto.UserLoginDTO)

		if err := c.BodyParser(data); err != nil {
			return SimpleError(c, err)
		}
		v := validator.New()

		err := v.Struct(data)
		if err != nil {
			return SimpleError(c, err)
		}

		rdb := sc.Redis
		db := sc.DB
		repo := repository.NewRepo(db)
		rdbStore := redix.NewRedisStore(rdb)
		userDomain := domain.NewUserDomain(repo, rdbStore)

		user, err := userDomain.Login(c.UserContext(), data.Username, data.Password)
		if err != nil {
			return SimpleError(c, err)
		}

		token, err := utils.GenerateJWT(user, sc)
		if err != nil {
			return SimpleError(c, err)
		}

		return c.Status(200).JSON(&fiber.Map{
			"token": token,
		})
	}
}

func UserSignup(sc *config.ServerConfig) fiber.Handler {
	return func(c *fiber.Ctx) error {
		data := new(dto.CreateUserDTO)

		if err := c.BodyParser(data); err != nil {
			return SimpleError(c, err)
		}

		if err := validator.New().Struct(data); err != nil {
			return SimpleError(c, err)
		}

		rdb := sc.Redis
		db := sc.DB
		repo := repository.NewRepo(db)
		rdbStore := redix.NewRedisStore(rdb)
		userDomain := domain.NewUserDomain(repo, rdbStore)

		user, err := userDomain.CreateUser(c.UserContext(), data.Username, data.Password)
		if err != nil {
			return SimpleError(c, err)
		}

		token, err := utils.GenerateJWT(user, sc)
		if err != nil {
			return SimpleError(c, err)
		}

		return c.Status(200).JSON(&fiber.Map{
			"token": token,
		})
	}
}
