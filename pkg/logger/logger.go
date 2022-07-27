package logger

import (
	"go.uber.org/zap"
)

var L *zap.Logger

func NewLogger() {
	logger, _ := zap.NewProduction()
	defer logger.Sync()
	L = logger
}
