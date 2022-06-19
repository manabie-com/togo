package task

import (
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"

	"togo/cache"
)

//go:generate mockery -name=Repository
type Repository interface {
	GetUserConfig(userId string) (*UserConfig, error)
	InsertUserTask(userId, task string, updatedAt time.Time) error
}

type Service struct {
	repo Repository
	cache cache.Redis
}

func NewService(repo Repository, cache cache.Redis) *Service {
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
	if userConfig.Limit == cachedLimit {
		return fmt.Errorf("user task record reached limit: %d", cachedLimit)
	}

	loc, _ := time.LoadLocation("Local")
	now := time.Now().In(loc)
	midnight := time.Date(now.Year(), now.Month(), now.Day(), 23, 59, 59, 99, loc)
	duration := midnight.Sub(now)
	err = s.repo.InsertUserTask(userId, task, now)
	if err != nil {
		return fmt.Errorf("insert user's task error: %v", err)
	}
	err = s.cache.SetExpire(userId, cachedLimit+1, duration)
	if err != nil {
		return fmt.Errorf("set redis expire error: %v", err)
	}
	return nil
}
