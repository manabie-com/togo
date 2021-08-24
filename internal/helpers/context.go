package helpers

import (
	"context"
	"github.com/manabie-com/togo/internal/constants"
)

func GetUserIdFromContext(ctx context.Context) (string, bool) {
	data := ctx.Value(constants.KeyUserId)
	userId, ok := data.(string)
	return userId, ok
}
