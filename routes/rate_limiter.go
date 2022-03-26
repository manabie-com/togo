package routes

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/gofiber/fiber/v2"
	"github.com/luongdn/togo/database"
	"github.com/luongdn/togo/models"
	"github.com/luongdn/togo/storage"
)

func RateLimiter(action models.UserAction) func(*fiber.Ctx) error {
	return func(ctx *fiber.Ctx) error {
		user_id := ctx.Params("user_id")
		store := storage.GetRuleStore()
		rule, _ := store.GetRule(ctx.Context(), user_id, action)

		if throttled(ctx.Context(), database.Rdb, rule) {
			log.Println("Throttled:", user_id, action)
			return ctx.SendStatus(429)
		}
		return ctx.Next()
	}
}

func throttled(ctx context.Context, rdb *redis.Client, rule models.Rule) bool {
	if rule.ID == "" {
		return false
	}

	counterKey := buildCounterKey(rule)
	count, err := rdb.Incr(ctx, counterKey).Result()
	if err != nil {
		panic(err)
	}
	if count == 1 {
		rdb.Expire(ctx, counterKey, getDuration(rule.Unit))
	}

	return count > rule.RequestsPerUnit
}

func buildCounterKey(rule models.Rule) string {
	window := getCurrentTimeWindow(rule.Unit)
	return fmt.Sprintf("counter:%s:%s:%s:%d", rule.UserID, rule.Action, rule.Unit, window)
}

func getDuration(unit models.TimeUnit) time.Duration {
	switch unit {
	case models.Second:
		return time.Second
	case models.Minute:
		return time.Minute
	case models.Hour:
		return time.Hour
	case models.Day:
		return time.Hour * 24
	default:
		return time.Hour * 24
	}
}

func getCurrentTimeWindow(unit models.TimeUnit) int {
	switch unit {
	case models.Second:
		return time.Now().Second()
	case models.Minute:
		return time.Now().Minute()
	case models.Hour:
		return time.Now().Hour()
	case models.Day:
		return time.Now().Day()
	default:
		return time.Now().Day()
	}
}
