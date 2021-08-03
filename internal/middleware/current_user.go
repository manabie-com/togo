package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
	"togo/config"
	"togo/internal/domain"
	"togo/internal/handler"
	"togo/internal/redix"
	"togo/internal/repository"
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
		hdl := handler.NewUserHandler(repo, rdbStore)
		userDomain := domain.NewUserDomain(hdl)

		currentUser, err := userDomain.GetUser(c.UserContext(), int32(userId))
		if err != nil {
			return err
		}

		c.Locals("currentUser", currentUser)

		return c.Next()
	}
}
