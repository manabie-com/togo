package rest

import (
	"errors"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
	"togo/config"
	"togo/internal/repository"
	"togo/internal/service"
)

type UserLoginDTO struct {
	Username string `json:"username" validate:"required,min=6,max=32"`
	Password string `json:"password" validate:"required,min=6,max=32"`
}

func UserLogin(sc *config.ServerConfig) fiber.Handler {
	return func(c *fiber.Ctx) error {
		data := new(UserLoginDTO)

		if err := c.BodyParser(data); err != nil {
			return SimpleError(c, err)
		}
		v := validator.New()

		err := v.Struct(data)
		if err != nil {
			return SimpleError(c, err)
		}

		db := sc.DB
		repo := repository.NewRepo(db)
		svc := service.NewUserService(repo)

		user, err := svc.GetUserByUsername(c.UserContext(), data.Username)
		if err != nil {
			return SimpleError(c, err)
		}
		userPass := []byte(data.Password)
		dbPass := []byte(user.Password)

		passErr := bcrypt.CompareHashAndPassword(dbPass, userPass)

		if passErr != nil {
			return SimpleError(c, errors.New("Wrong password"))
		}

		token, err := GenerateJWT(user, sc)
		if err != nil {
			return SimpleError(c, err)
		}

		return c.Status(200).JSON(&fiber.Map{
			"token": token,
		})
	}
}

type CreateUserDTO struct {
	Username string `json:"username" validate:"required,min=6,max=32"`
	Password string `json:"password" validate:"required,min=6,max=32"`
}

func UserSignup(sc *config.ServerConfig) fiber.Handler {
	return func(c *fiber.Ctx) error {
		data := new(CreateUserDTO)

		if err := c.BodyParser(data); err != nil {
			return SimpleError(c, err)
		}
		v := validator.New()

		err := v.Struct(data)
		if err != nil {
			return SimpleError(c, err)
		}

		db := sc.DB
		repo := repository.NewRepo(db)
		svc := service.NewUserService(repo)

		data.Password = GetHash([]byte(data.Password))

		user, err := svc.CreateUser(c.UserContext(), data.Username, data.Password)
		if err != nil {
			return SimpleError(c, err)
		}

		token, err := GenerateJWT(&user, sc)
		if err != nil {
			return SimpleError(c, err)
		}

		return c.Status(200).JSON(&fiber.Map{
			"token": token,
		})
	}
}
