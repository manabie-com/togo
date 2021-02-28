package helper

import (
	"sync"
	"time"
)

type ToDos struct {
	current   int
	max       int
	createdAt time.Time

	mu sync.Mutex
}

func NewToDos(max, current int) *ToDos {
	return &ToDos{
		current:   current,
		max:       max,
		createdAt: time.Now(),
	}
}

func (t *ToDos) CanAddNewTodo() bool {
	t.mu.Lock()
	defer t.mu.Unlock()

	now := time.Now()
	year, month, day := t.createdAt.Date()
	yearNow, monthNow, dayNow := now.Date()

	// when date(now) <> date(createdAt) then createdAt = now and reset current to zero
	if year != yearNow || month != monthNow || day != dayNow {
		t.createdAt = now
		t.current = 0
	}

	if t.current < t.max {
		t.current += 1
		return true
	}

	return false
}
