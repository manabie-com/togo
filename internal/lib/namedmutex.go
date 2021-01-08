package lib

import "sync"

type NamedMutex struct {
	pool  sync.Pool
	locks sync.Map
}

func NewNamedMutex() *NamedMutex {
	return &NamedMutex{
		pool: sync.Pool{
			New: func() interface{} {
				return new(sync.Mutex)
			},
		},
	}
}

func (nm *NamedMutex) Lock(name string) {
	nm.getLock(name).Lock()
}

func (nm *NamedMutex) Unlock(name string) {
	nm.getLock(name).Unlock()
}

func (nm *NamedMutex) getLock(name string) *sync.Mutex {
	newLock := nm.pool.Get()
	lock, stored := nm.locks.LoadOrStore(name, newLock)
	if stored {
		nm.pool.Put(newLock)
	}
	return lock.(*sync.Mutex)
}
