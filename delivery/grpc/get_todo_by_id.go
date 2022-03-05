package grpc

import (
	"context"
	"encoding/json"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/khangjig/togo/proto"
	"github.com/khangjig/togo/util/myerror"
)

func (t *TogoService) GetTodoByID(ctx context.Context, req *proto.GetByIDRequest) (*proto.Todo, error) {
	myTodo, err := t.UseCase.GRPC.GetTodoByID(ctx, req.Id)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, err.(myerror.MyError).Message)
	}

	resp := &proto.Todo{}

	b, _ := json.Marshal(myTodo)
	_ = json.Unmarshal(b, resp)

	return resp, nil
}
