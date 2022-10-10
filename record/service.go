package record

import (
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)



//go:generate mockery -name=Repository
type Repository interface {
	GetUserConfig(userId string) (*UserConfig, error)
	InsertUserTask(userId, task string, updatedAt time.Time) error
}

//go:generate mockery -name=CacheService
type CacheService interface {
	GetInt(key string) (int, error)
	SetExpire(key string, value int, expire time.Duration) error
}

type Service struct {
	repo Repository
	cache CacheService
}

func NewService(repo Repository, cache CacheService) *Service {
	return &Service{
		repo: repo,
		cache: cache,
	}
}

func (s *Service) RecordTask(userId, task string) error {
	userConfig, err := s.repo.GetUserConfig(userId)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return fmt.Errorf("user id does not exist: %s", userId)
		}
		return fmt.Errorf("get user config error: %v", err)
	}

	cachedLimit, err := s.cache.GetInt(userId)
	if err != nil {
		return fmt.Errorf("get redis error: %v", err)
	}
	if cachedLimit >= userConfig.Limit {
		return fmt.Errorf("user record record reached limit: %d", cachedLimit)
	}

	loc, _ := time.LoadLocation("Local")
	now := time.Now().In(loc)
	midnight := time.Date(now.Year(), now.Month(), now.Day(), 23, 59, 59, 999999999, loc)
	duration := midnight.Sub(now)
	err = s.repo.InsertUserTask(userId, task, now)
	if err != nil {
		return fmt.Errorf("insert user's record error: %v", err)
	}
	err = s.cache.SetExpire(userId, cachedLimit+1, duration)
	if err != nil {
		return fmt.Errorf("set redis expire error: %v", err)
	}
	return nil
}
