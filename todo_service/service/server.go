package service

import (
	"context"
	"database/sql"

	"todo_service/proto"
)

//serviceServer ...
type serviceServer struct {
	db *sql.DB
	proto.UnimplementedTodoServiceServer
}

//NewTodoServiceServer ...
func NewTodoServiceServer(db *sql.DB) proto.TodoServiceServer {
	return &serviceServer{db: db}
}

func (s *serviceServer) CreateTodo(ctx context.Context, request *proto.CreateTodoRequest) (*proto.CreateTodoResponse, error) {
	return nil, nil
}

func (s *serviceServer) GetTodo(ctx context.Context, request *proto.GetTodoRequest) (*proto.GetTodoResponse, error) {
	return nil, nil
}

func (s *serviceServer) UpdateTodo(ctx context.Context, request *proto.UpdateTodoRequest) (*proto.UpdateTodoResponse, error) {
	return nil, nil
}

func (s *serviceServer) DeleteTodo(ctx context.Context, request *proto.DeleteTodoRequest) (*proto.DeleteTodoResponse, error) {
	return nil, nil
}
