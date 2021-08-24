package tools

import (
	"context"
	"net/http"
)

type UserAuthKey int8

type IContextTool interface {
	UserIDFromCtx(ctx context.Context) (string, *TodoError)
	WriteUserIDToContext(ctx context.Context, id string) context.Context
}

type ContextTool struct{}

func (ct *ContextTool) UserIDFromCtx(ctx context.Context) (string, *TodoError) {
	v := ctx.Value(UserAuthKey(0))
	id, ok := v.(string)
	if !ok {
		return "", NewTodoError(http.StatusInternalServerError, "Something wrong when set user id")
	}
	return id, nil
}

func (ct *ContextTool) WriteUserIDToContext(ctx context.Context, id string) context.Context {
	return context.WithValue(ctx, UserAuthKey(0), id)
}

func NewContextTool() IContextTool {
	return &ContextTool{}
}
