package middleware

import (
	"time"

	"github.com/dgrijalva/jwt-go"

	"github.com/manabie-com/togo/internal/config"
	"github.com/manabie-com/togo/internal/models"
)

func CreateAccountJWT(account models.Account) (string, error) {
	//create jwt
	claims := jwt.MapClaims{}
	claims["username"] = account.Username
	claims["account_id"] = account.ID
	claims["max_daily_tasks_count"] = account.MaxDailyTasksCount
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()
	tmp := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tmp.SignedString([]byte(config.GetConfig().JWTSecret))
	return token, err
}
