package utils

import (
	"context"
	"errors"
	"github.com/manabie-com/togo/internal/api/dictionary"
)

func GetValueFromCtx(ctx context.Context, field string) (string, error) {
	v := ctx.Value(field)
	id, ok := v.(string)
	if !ok {
		return "", errors.New(dictionary.FailedToGetValueFromContext)
	}
	return id, nil
}
