package ratelimiters

import (
	"fmt"
	"sync"
	"time"

	sqllite "github.com/manabie-com/togo/internal/storages/sqlite"
)

type elementRateLimter struct {
	Mutex     sync.RWMutex
	Count     int
	ExpiredAt time.Time
}

type localRateLimiter struct {
}

var (
	// For multi instance api servers, should use redis instead.
	localRL                 = map[string]*elementRateLimter{}
	_       baseRateLimiter = localRateLimiter{}
)

func InitLocalRatelimiter(s *sqllite.LiteDB) {
	if ratelimiter != nil {
		panic(fmt.Errorf("Duplicate ratelimiter"))
	}

	totalTasks := s.GetTotalTaskToDayByAllUser()
	for userId, totalTask := range totalTasks {
		localRL[GenTaskTarget(userId)] = &elementRateLimter{
			Mutex:     sync.RWMutex{},
			Count:     totalTask,
			ExpiredAt: GenTaskLimitExpiredAt(),
		}
	}
	ratelimiter = localRateLimiter{}
}

func (r localRateLimiter) Increase(target string) int {
	count := 0
	localRL[target].Mutex.Lock()
	if localRL[target].ExpiredAt.Before(time.Now()) {
		localRL[target].ExpiredAt = GenTaskLimitExpiredAt()
		localRL[target].Count = 0
	}
	localRL[target].Count++
	count = localRL[target].Count
	localRL[target].Mutex.Unlock()
	return count
}

func (r localRateLimiter) Decrease(target string) {
	localRL[target].Mutex.Lock()
	localRL[target].Count--
	localRL[target].Mutex.Unlock()
}
