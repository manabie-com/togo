package server

import (
	"context"
	db_api "mini_project/db"
	"mini_project/db/model"
	rpc_api "mini_project/rpc_services"
	"sync"
)

// define event on the system for pubsub pattern with redis
type event string

// implements the APIServer
type APIServer struct {
	db model.DatabaseAPI
	m  *sync.RWMutex
}

// var policy embed.FS

func NewAPIServer(dbUrl map[string]string) (*APIServer, error) {

	// connect mysql
	db, err := db_api.GetDatabase(dbUrl)

	if err != nil {
		return nil, err
	}

	if err != nil {
		return nil, err
	}

	return &APIServer{db: db, m: &sync.RWMutex{}}, err
}

//NewTask(context.Context, *NewTaskReq) (*NewTaskResp, error)
func (s *APIServer) NewTask(ctx context.Context, req *rpc_api.NewTaskReq) (*rpc_api.NewTaskResp, error) {

	return nil, nil
}
