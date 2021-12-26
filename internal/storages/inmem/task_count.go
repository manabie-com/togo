package inmem

import (
	"strings"
	"sync"
	"time"

	"github.com/perfectbuii/togo/internal/storages"
)

type taskCountStore struct {
	lock *sync.Mutex
	data map[string]int
}

func (s *taskCountStore) Value(key string) int {
	defer s.lock.Unlock()
	s.lock.Lock()
	v, ok := s.data[key]
	if !ok {
		s.data[key] = 0
		return 0
	}
	return v
}

func (s *taskCountStore) Inc(key string) int {
	defer s.lock.Unlock()
	s.lock.Lock()
	if v, ok := s.data[key]; ok {
		s.data[key] = v + 1
		return s.data[key]
	} else {
		s.data[key] = 1
		return 1
	}
}

func (s *taskCountStore) Desc(key string) {
	defer s.lock.Unlock()
	s.lock.Lock()
	if v, ok := s.data[key]; ok {
		s.data[key] = v - 1
	} else {
		s.data[key] = 0
	}
}

func (s *taskCountStore) gc() {
	for {
		currentDate := time.Now().Format("2006-01-02")
		time.Sleep(24 * time.Hour)
		for i := range s.data {
			if strings.Contains(i, currentDate) {
				delete(s.data, i)
			}
		}
	}
}

func NewTaskCountStore() storages.TaskCountStore {
	t := &taskCountStore{
		lock: &sync.Mutex{},
		data: map[string]int{},
	}

	go t.gc()
	return t
}
