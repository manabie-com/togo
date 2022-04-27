package user

import (
	"fmt"
	"time"
	"togo/internal/model"
	"togo/internal/services/login"
	"togo/internal/utils"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
)

// Login get user and password
func Login(c *fiber.Ctx) error {

	payload := new(model.LoginInput)
	utils.BodyParser(c, payload)

	user := payload.Username
	pass := payload.Password

	fmt.Println("user: ", user)
	fmt.Println("pass: ", pass)

	if !login.Validate(user, pass) {
		return c.JSON(fiber.Map{"status": "error", "message": "Invalid username/password"})
	}

	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["identity"] = user
	claims["admin"] = true
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.JSON(fiber.Map{"status": "success", "message": "Success login", "data": t})
}
