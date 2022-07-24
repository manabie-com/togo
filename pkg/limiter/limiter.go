package limiter

import (
	"sync"
	"time"

	sw "github.com/RussellLuo/slidingwindow"
)

var DefaultLimit = Limit{
	size:  24 * time.Hour,
	limit: 5,
}

type (
	Limit struct {
		size  time.Duration
		limit int64
	}
	Limiter struct {
		limitByID map[string]Limit
		limiters  map[string]*sw.Limiter

		locker sync.Mutex
	}
)

func New() *Limiter {
	return &Limiter{
		limitByID: make(map[string]Limit),
		limiters:  make(map[string]*sw.Limiter, 0),
	}
}

func (l *Limiter) WithLimitByID(id string, size time.Duration, limit int64) {
	l.limitByID[id] = Limit{size: size, limit: limit}
	l.WithLimiter(id, l.Limiter(id))
}

func (l *Limiter) Limit(id string) Limit {
	if lim, found := l.limitByID[id]; found {
		return lim
	}

	return DefaultLimit
}

func (l *Limiter) WithLimiter(id string, lim *sw.Limiter) {
	l.limiters[id] = lim
}

func (l *Limiter) Limiter(id string) *sw.Limiter {
	if lim, found := l.limiters[id]; found {
		return lim
	}

	lim := l.Limit(id)
	limiter, _ := sw.NewLimiter(lim.size, lim.limit, func() (sw.Window, sw.StopFunc) {
		return sw.NewLocalWindow()
	})
	l.WithLimiter(id, limiter)

	return limiter
}

func (l *Limiter) AllowN(t time.Time, id string, n int64) bool {
	l.locker.Lock()
	defer l.locker.Unlock()

	return l.Limiter(id).AllowN(t, n)
}
