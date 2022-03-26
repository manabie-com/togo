package storage

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/luongdn/togo/models"
	"gorm.io/gorm"
)

func NewRuleStore(db *gorm.DB, rdb *redis.Client) *ruleStore {
	return &ruleStore{
		sqlDB: db,
		rdb:   rdb,
	}
}

type ruleStore struct {
	sqlDB *gorm.DB
	rdb   *redis.Client
}

func (s *ruleStore) GetRule(ctx context.Context, user_id string, action models.UserAction) (models.Rule, error) {
	rule, _ := s.GetCache(ctx, user_id, action)
	// Cache hit
	if rule.ID != "" {
		return rule, nil
	}

	rule, _ = s.GetFromDB(ctx, user_id, action)
	s.setCache(ctx, rule)
	return rule, nil
}

func buildRuleKey(user_id string, action models.UserAction) string {
	return fmt.Sprintf("rule:%s:%s", user_id, action)
}

func (s *ruleStore) GetCache(ctx context.Context, user_id string, action models.UserAction) (models.Rule, error) {
	key := buildRuleKey(user_id, action)
	val, err := s.rdb.Get(ctx, key).Result()

	if err != nil {
		return models.Rule{}, err
	}

	rule := models.Rule{}
	json.Unmarshal([]byte(val), &rule)
	return rule, nil
}

func (s *ruleStore) GetFromDB(ctx context.Context, user_id string, action models.UserAction) (models.Rule, error) {
	rule := models.Rule{}
	result := s.sqlDB.Where("user_id = ? AND action = ?", user_id, action).First(&rule)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return models.Rule{}, nil
	}
	if result.Error != nil {
		return models.Rule{}, result.Error
	}
	return rule, nil
}

func (s *ruleStore) setCache(ctx context.Context, rule models.Rule) {
	key := buildRuleKey(rule.UserID, rule.Action)
	val, err := json.Marshal(rule)
	if err != nil {
		panic(err)
	}
	s.rdb.Set(ctx, key, val, time.Hour*24)
}

func (s *ruleStore) delCache(ctx context.Context, user_id string, action models.UserAction) {
	key := buildRuleKey(user_id, action)
	s.rdb.Del(ctx, key)
}

func (s *ruleStore) CreateRule(ctx context.Context, user_id string, rule *models.Rule) error {
	rule.UserID = user_id
	result := s.sqlDB.Create(rule)

	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (s *ruleStore) UpdateRule(ctx context.Context, user_id string, rule *models.Rule) error {
	rule.UserID = user_id
	result := s.sqlDB.Model(&models.Rule{}).Where("user_id = ? AND action = ?", rule.UserID, rule.Action).Updates(rule)

	if result.Error != nil {
		return result.Error
	}
	s.delCache(context.Background(), rule.UserID, rule.Action)
	return nil
}
