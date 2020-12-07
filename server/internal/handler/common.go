package handler

import (
	"context"
	"github.com/HoangVyDuong/togo/pkg/define"
)

func UserIDFromContext(ctx context.Context) (uint64, error) {
	userID := ctx.Value(define.ContextKeyUserID)
	numUserID, ok := userID.(uint64)
	if !ok {
		return 0, define.FailedValidation
	}
	if numUserID == 0 {
		return 0, define.FailedValidation
	}
	return numUserID, nil
}
