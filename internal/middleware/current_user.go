package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
	"togo/config"
	"togo/internal/redix"
	"togo/internal/repository"
	"togo/internal/domain"
)

func SetCurrentUser(sc *config.ServerConfig) fiber.Handler {
	return func(c *fiber.Ctx) error {
		user := c.Locals("user").(*jwt.Token)
		claims := user.Claims.(jwt.MapClaims)
		userId := claims["user_id"].(float64)

		rdb := sc.Redis
		db := sc.DB
		repo := repository.NewRepo(db)
		rdbStore := redix.NewRedisStore(rdb)
		svc := domain.NewUserDomain(repo, rdbStore)

		currentUser, err := svc.GetUser(c.UserContext(), int32(userId))
		if err != nil {
			return err
		}

		c.Locals("currentUser", currentUser)

		return c.Next()
	}
}
