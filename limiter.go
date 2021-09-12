package togo

import (
	"golang.org/x/time/rate"
)

type UserLimiter interface {
	AddUser(userID string) *rate.Limiter
	GetLimiter(userID string) *rate.Limiter
}
