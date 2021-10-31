package auth

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/pkg/errors"
	"github.com/quochungphp/go-test-assignment/src/pkgs/redis"
	"github.com/quochungphp/go-test-assignment/src/pkgs/settings"
	"github.com/quochungphp/go-test-assignment/src/pkgs/token"
)

// AuthLogoutAction ...
type AuthLogoutAction struct{}

// Execute ...
func (Auth AuthLogoutAction) Execute() (err error) {
	sessionUser := token.AccessUser

	if sessionUser.UserID == 0 {
		return errors.Errorf("Invalid session")
	}

	expiresIn, err := strconv.Atoi(os.Getenv(settings.RedisCacheExpiresIn))
	if err != nil {
		return errors.Wrap(err, "Error while parsing redis expires in")
	}

	cacheKey := redis.TokenBlackListCacheKey(fmt.Sprintf("%s-%d", sessionUser.CorrelationID, sessionUser.UserID))
	err = redis.SetItem(cacheKey, true, time.Duration(expiresIn)*time.Second)
	if err != nil {
		return errors.Wrap(err, "Error while creating redis expires in")
	}

	return nil
}
