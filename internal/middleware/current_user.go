package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
	"togo/config"
	"togo/internal/repository"
	"togo/internal/service"
)

func SetCurrentUser(sc *config.ServerConfig) fiber.Handler {
	return func(c *fiber.Ctx) error {
		user := c.Locals("user").(*jwt.Token)
		claims := user.Claims.(jwt.MapClaims)
		userId := claims["user_id"].(float64)

		db := sc.DB
		repo := repository.NewRepo(db)
		svc := service.NewUserService(repo)

		currentUser, err := svc.GetUser(c.UserContext(), int32(userId))
		if err != nil {
			return err
		}

		c.Locals("currentUser", currentUser)

		return c.Next()
	}
}
