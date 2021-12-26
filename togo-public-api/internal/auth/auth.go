package auth

import (
	"context"
	"strings"
	"togo-public-api/internal/service/togo_user_session_v1"

	"github.com/giahuyng98/togo/core-lib/logger"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

type authKeyType struct{}

var (
	ErrUnauthenticated = status.Error(codes.Unauthenticated, codes.Unauthenticated.String())
	ErrInternal        = status.Error(codes.Internal, "Internal Server Error")
)

func contains(list []string, value string) bool {
	for _, item := range list {
		if item == value {
			return true
		}
	}
	return false
}

func NewAuthFunc(session togo_user_session_v1.TogoUserSessionServiceClient, whitelistMethods ...string) func(context.Context) (context.Context, error) {
	return func(ctx context.Context) (context.Context, error) {

		// Skip white list methods
		if method, ok := grpc.Method(ctx); ok {
			shortMethod := method[strings.LastIndex(method, "/")+1:]
			if contains(whitelistMethods, method) || contains(whitelistMethods, shortMethod) {
				return ctx, nil
			}
		}

		md, ok := metadata.FromIncomingContext(ctx)
		if !ok {
			return ctx, ErrUnauthenticated
		}
		tokens, ok := md["authorization"]
		if !ok || len(tokens) == 0 {
			return ctx, ErrUnauthenticated
		}
		token := strings.TrimPrefix(tokens[0], "Bearer ")
		if len(token) == 0 {
			return ctx, ErrUnauthenticated
		}
		resp, err := session.VerifyToken(ctx, &togo_user_session_v1.VerifyTokenRequest{
			Token: token,
		})
		if err != nil {
			logger.For(ctx).Error("VerifyToken error", zap.Error(err))
			return ctx, ErrInternal
		}
		ctx = context.WithValue(ctx, authKeyType{}, resp.UserId)
		return ctx, nil
	}
}

func GetUserID(ctx context.Context) string {
	if val, ok := ctx.Value(authKeyType{}).(string); ok {
		return val
	}
	return ""
}
