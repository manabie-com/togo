package service

import (
	"context"
	"database/sql"

	"account_service/connection"
	"account_service/proto"
	"account_service/utils"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// var clientSessionService proto.SessionServiceClient

// //DialToServiceServer ...
// func DialToServiceServer(port string) {

// 	conn, err := grpc.Dial("localhost:"+port, grpc.WithInsecure())
// 	if err != nil {
// 		panic(err)
// 	}
// 	clientSessionService = proto.NewSessionServiceClient(conn)
// }

//serviceServer ...
type serviceServer struct {
	db *sql.DB
	proto.UnimplementedAccountServiceServer
}

//NewAccountServiceServer ...
func NewAccountServiceServer(db *sql.DB) proto.AccountServiceServer {
	return &serviceServer{db: db}
}

func (s *serviceServer) Create(ctx context.Context, request *proto.CreateRequest) (*proto.CreateResponse, error) {

	conn, err := s.db.Conn(ctx)
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	hashedPassword := utils.HashAndSalt(request.GetPassword())
	res, err := conn.ExecContext(ctx, "INSERT INTO Account (Name, Email, Username, Password) VALUES (?, ?, ?, ?)",
		request.GetName(), request.GetEmail(), request.GetUsername(), hashedPassword)

	if err != nil {
		return nil, status.Error(codes.Unknown, "Failed to insert into Account-> "+err.Error())
	}

	id, err := res.LastInsertId()
	if err != nil {
		return nil, status.Error(codes.Unknown, "Failed to retrieve id for created Account-> "+err.Error())
	}
	_ = id

	return &proto.CreateResponse{
		IsCreated: true,
	}, nil
}

func (s *serviceServer) Login(ctx context.Context, request *proto.LoginRequest) (*proto.LoginResponse, error) {

	conn, err := s.db.Conn(ctx)
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	rows, err := conn.QueryContext(ctx, "SELECT Password, Id FROM Account WHERE Username = ?", request.GetUsername())
	if err != nil {
		return nil, status.Error(codes.Unknown, "Failed to select Username from Account-> "+err.Error())
	}
	defer rows.Close()

	var hashedPassoword string
	var id uint64

	for rows.Next() {
		err = rows.Scan(&hashedPassoword, &id)
		if err != nil {
			return nil, err
		}
	}

	pwdErr := utils.ComparePasswords(hashedPassoword, request.GetPassword())
	if pwdErr != nil {
		return nil, pwdErr
	}

	sc := connection.DialToSessionServiceServer()
	tokenString, err := sc.ClientSessionService.CreateToken(ctx, &proto.AccountInfo{
		Id: id,
	})

	return &proto.LoginResponse{
		Token: tokenString.Token,
	}, nil
}

func (s *serviceServer) Logout(ctx context.Context, request *proto.LogoutRequest) (*proto.LogoutResponse, error) {

	sc := connection.DialToSessionServiceServer()
	stat, err := sc.ClientSessionService.DeleteToken(ctx, &proto.TokenString{
		Token: request.GetToken(),
	})
	if err != nil {
		return nil, err
	}

	return &proto.LogoutResponse{
		IsLoggedOut: stat.Success,
	}, nil
}

func (s *serviceServer) Update(ctx context.Context, request *proto.UpdateRequest) (*proto.UpdateResponse, error) {
	return nil, nil
}

func (s *serviceServer) Show(ctx context.Context, request *proto.ShowRequest) (*proto.ShowResponse, error) {
	return nil, nil
}
