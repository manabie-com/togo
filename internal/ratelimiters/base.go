package ratelimiters

import (
	"fmt"
	"time"
)

var (
	ratelimiter baseRateLimiter
)

type baseRateLimiter interface {
	Increase(target string) int
	Decrease(target string)
}

func Increase(target string) int {
	return ratelimiter.Increase(target)
}

func Decrease(target string) {
	ratelimiter.Decrease(target)
}

func GenTaskTarget(userId string) string {
	return fmt.Sprintf("Task-%s", userId)
}

func GenTaskLimitExpiredAt() time.Time {
	return time.Now().Round(24*time.Hour).AddDate(0, 0, 1)
}
