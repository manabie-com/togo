package services

import (
	"sync"

	"golang.org/x/time/rate"
)

type userRateLimiter struct {
	users map[string]*rate.Limiter
	mu    *sync.RWMutex
	r rate.Limit
	b int
}

func NewUserRateLimiter(r rate.Limit, b int) *userRateLimiter {
	return &userRateLimiter{
		users: make(map[string]*rate.Limiter),
		mu:    &sync.RWMutex{},
		r:     r,
		b:     b,
	}
}

func (u *userRateLimiter) AddUser(userID string) *rate.Limiter {
	u.mu.Lock()
	defer u.mu.Unlock()

	limiter := rate.NewLimiter(u.r, u.b)

	u.users[userID] = limiter
	return limiter
}

func (u *userRateLimiter) GetLimiter(userID string) *rate.Limiter {
	u.mu.Lock()

	limiter, exists := u.users[userID]
	if !exists {
		u.mu.Unlock()
		return u.AddUser(userID)
	}

	u.mu.Unlock()

	return limiter
}
