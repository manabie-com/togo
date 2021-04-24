package psql

import "sync"

type simpleLock struct {
	internal *sync.Mutex
}

func (s simpleLock) NewMutex(string) Mutex {
	return simpleMutex{internal: s.internal}
}

type simpleMutex struct {
	internal *sync.Mutex
}

func (s simpleMutex) Lock() error {
	s.internal.Lock()
	return nil
}

func (s simpleMutex) Unlock() error {
	s.internal.Unlock()
	return nil
}

type Lock interface {
	NewMutex(session string) Mutex
}

type Mutex interface {
	Lock() error
	Unlock() error
}
