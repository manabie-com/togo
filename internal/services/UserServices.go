package services

import (
	"time"
	"togo/internal/dao"
	"togo/internal/model"
	"togo/internal/utils"

	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
)

type UserServices interface {
	LoginUser() error
}

func (c Con) LoginUser() error {

	payload := new(model.LoginInput)
	utils.BodyParser(c.Ctx, payload)

	user := payload.Username
	pass := payload.Password

	if !dao.CheckCredential(user, pass, c.Db) {
		return c.Ctx.JSON(fiber.Map{"status": "error", "message": "Invalid username/password"})
	}

	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["identity"] = user
	claims["admin"] = true
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		return c.Ctx.SendStatus(fiber.StatusInternalServerError)
	}

	return c.Ctx.JSON(fiber.Map{"status": "success", "message": "Success login", "data": t})
}

func (c Con) CreateUser() error {

	payload := new(model.User)
	utils.BodyParser(c.Ctx, payload)
	validate := validator.New()

	if err := validate.Struct(payload); err != nil {
		return c.Ctx.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	if !dao.CheckUserExist(payload.UserName, c.Db) {
		return c.Ctx.Status(fiber.StatusConflict).JSON(fiber.Map{"status": "error", "message": "Username already exists"})
	}

	hash, err := utils.HashPassword(payload.Password)
	if err != nil {
		return c.Ctx.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	payload.Password = hash

	result := c.Db.Create(&payload)
	if result.Error != nil {
		return c.Ctx.Status(500).JSON(fiber.Map{"retCode": 500, "message": result.Error})

	}

	return c.Ctx.JSON(payload)
}
