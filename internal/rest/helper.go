package rest

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
	"log"
	"time"
	"togo/config"
	"togo/internal/entity"
	"togo/utils/validator"
)

const appErrFormErrResponseFailure = "form error response failure"

func SimpleError(c *fiber.Ctx, err error) error {
	resp := validator.ToErrResponse(err)

	if resp == nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": fiber.Map{"errors": appErrFormErrResponseFailure},
		})
	}

	return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
		"message": resp,
	})
}

func GenerateJWT(user *entity.User, sc *config.ServerConfig) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["user_id"] = user.ID
	claims["user_name"] = user.Username
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	return token.SignedString([]byte(sc.JwtSecret))
}

func GetHash(pwd []byte) string {
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil {
		log.Println(err)
	}
	return string(hash)
}

func GetCurrentUserId(c *fiber.Ctx) int32 {
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["user_id"].(float64)

	return int32(userId)
}
