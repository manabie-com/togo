package util

import "context"

const UserAuthKey = int8(0)

func SetUserIDToContext(ctx context.Context, userID string) context.Context {
	return context.WithValue(ctx, UserAuthKey, userID)
}

func GetUserIDFromContext(ctx context.Context) (string, bool) {
	v := ctx.Value(UserAuthKey)
	id, ok := v.(string)
	return id, ok
}
