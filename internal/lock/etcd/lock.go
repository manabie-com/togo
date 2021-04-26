package etcd

import (
	"context"

	"github.com/manabie-com/togo/internal/lock"
	"go.etcd.io/etcd/clientv3"
	"go.etcd.io/etcd/clientv3/concurrency"
)

type Config struct {
	Endpoints []string
}

type locker struct {
	client *clientv3.Client
}

type session struct {
	s   *concurrency.Session
	mut *concurrency.Mutex
}

func NewLock(c Config) (*locker, error) {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints: c.Endpoints,
	})
	if err != nil {
		return nil, err
	}
	return &locker{
		client: cli,
	}, nil
}

func (l *locker) NewMutex(key string) (lock.Mutex, error) {
	s, err := concurrency.NewSession(l.client)
	if err != nil {
		return nil, err
	}

	mut := concurrency.NewMutex(s, key)
	return &session{
		s:   s,
		mut: mut,
	}, nil
}

func (s *session) Lock() error {
	err := s.mut.Lock(context.Background())
	if err != nil {
		s.s.Close()
		return err
	}
	return nil
}

func (s *session) Unlock() error {
	defer s.s.Close()
	err := s.mut.Unlock(context.Background())
	if err != nil {
		return err
	}
	return nil
}
