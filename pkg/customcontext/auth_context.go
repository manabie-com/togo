package customcontext

import "context"

const userAuthKey = ctxKey(0)

type ctxKey int8

func UserIDFromCtx(ctx context.Context) string {
	v := ctx.Value(userAuthKey)
	userID, _ := v.(string)
	return userID
}

func SetUserIDToContext(ctx context.Context, userID string) context.Context {
	return context.WithValue(ctx, userAuthKey, userID)
}
