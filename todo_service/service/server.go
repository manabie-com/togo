package service

import (
	"context"
	"database/sql"

	"todo_service/connection"
	"todo_service/proto"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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

	sc := connection.DialToSessionServiceServer()
	stat, err := sc.ClientSessionService.CheckToken(ctx, &proto.TokenString{
		Token: request.GetToken(),
	})
	if err != nil {
		return nil, err
	}
	accountID, err := sc.ClientSessionService.GetAccountIDFromToken(ctx, &proto.TokenString{
		Token: request.GetToken(),
	})
	if err != nil {
		return nil, err
	}
	accId := &accountID.Id

	conn, err := s.db.Conn(ctx)
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	res, err := conn.ExecContext(ctx, "INSERT INTO Todo (Title, Description, Status, AccountId) VALUES (?, ?, ?, ?)",
		request.GetTitle(), request.GetDescription(), request.GetStatus(), accId)

	if err != nil {
		return nil, status.Error(codes.Unknown, "Failed to insert into Todo-> "+err.Error())
	}

	insertedId, err := res.LastInsertId()
	if err != nil {
		return nil, status.Error(codes.Unknown, "Failed to retrieve id for created Todo-> "+err.Error())
	}

	return &proto.CreateTodoResponse{
		Success: stat.Success,
		TodoId:  insertedId,
	}, nil
}

func (s *serviceServer) GetTodo(ctx context.Context, request *proto.GetTodoRequest) (*proto.GetTodoResponse, error) {

	conn, err := s.db.Conn(ctx)
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	rows, err := conn.QueryContext(ctx, "SELECT Title, Description, Status FROM Todo WHERE Id = ?", request.GetTodoId())
	if err != nil {
		return nil, status.Error(codes.Unknown, "Failed to select from Todo-> "+err.Error())
	}
	defer rows.Close()

	type todo struct {
		Title       sql.NullString
		Description sql.NullString
		Status      sql.NullInt32
	}

	var getTodoResponse todo
	for rows.Next() {
		err = rows.Scan(&getTodoResponse.Title, &getTodoResponse.Description, &getTodoResponse.Status)
		if err != nil {
			return nil, err
		}
	}

	return &proto.GetTodoResponse{
		Title:       getTodoResponse.Title.String,
		Description: getTodoResponse.Description.String,
		Status:      proto.TodoStatus(getTodoResponse.Status.Int32),
	}, nil
}

func (s *serviceServer) UpdateTodo(ctx context.Context, request *proto.UpdateTodoRequest) (*proto.UpdateTodoResponse, error) {
	return nil, nil
}

func (s *serviceServer) DeleteTodo(ctx context.Context, request *proto.DeleteTodoRequest) (*proto.DeleteTodoResponse, error) {
	return nil, nil
}
