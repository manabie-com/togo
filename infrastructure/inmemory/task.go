package inmemory

import (
	"context"
	"sync"
	"time"
	"togo/domain/errdef"
	"togo/domain/model"
	"togo/domain/repository"
)

type inmemoryTaskRepo struct {
	task map[string]model.Task
	mtx  sync.RWMutex
}

func (this *inmemoryTaskRepo) CountTaskCreatedInDayByUser(ctx context.Context, u model.User) (int, error) {
	this.mtx.Lock()
	defer this.mtx.Unlock()
	count := 0
	currentDate := time.Now().Format("2006-01-02")
	for _, t := range this.task {
		if t.CreatedBy == u.Id && t.CreatedDate == currentDate {
			count++
		}
	}
	return count, nil
}

func (this *inmemoryTaskRepo) Create(ctx context.Context, u model.Task) error {
	//TODO implement me
	this.mtx.Lock()
	defer this.mtx.Unlock()
	if _, ok := this.task[u.Title]; ok {
		return errdef.DupplicateTask
	}
	u.CreatedTime = time.Now()
	u.CreatedDate = time.Now().Format("2006-01-02")
	this.task[u.Title] = u
	return nil
}

func NewInMemoryTaskRepository() repository.TaskRepository {
	return &inmemoryTaskRepo{
		task: make(map[string]model.Task),
	}
}
