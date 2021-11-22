package server

import (
	"context"
	"fmt"
	"mini_project/common"
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

// AuthFuncOverride ()
func (s *APIServer) AuthFuncOverride(ctx context.Context, fullMethodName string) (context.Context, error) {
	return common.VerifyTokenBearer(ctx, fullMethodName, s.db)
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
	fmt.Println("<<<<<<<<<NewTask>>>>>>>>>")
	err := s.db.CreateTask(req.UserId, req.TaskName)
	rsp := rpc_api.NewTaskResp{TaskName: req.TaskName}
	return &rsp, err
}
