package transport

import (
	"sync"
	"time"
)

var (
	limiter = NewLimitStore()
)

type Limit struct {
	Count    int
	Interval time.Time
}
type LimitStore struct {
	mapIP          map[string]Limit
	intervalSecond float64
	maxRequest     int
	mu             *sync.RWMutex
}

func NewLimitStore() *LimitStore {
	return &LimitStore{
		mapIP:          make(map[string]Limit),
		intervalSecond: 2,
		maxRequest:     10,
		mu:             &sync.RWMutex{},
	}
}

func (l *LimitStore) GetIP(ip string) Limit {
	l.mu.Lock()
	defer l.mu.Unlock()

	limit, _ := l.mapIP[ip]

	return limit
}

func (l *LimitStore) AddIP(ip string, limit Limit) {
	l.mu.Lock()
	defer l.mu.Unlock()

	l.mapIP[ip] = limit
}
