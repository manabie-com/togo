package logger

import (
	"go.uber.org/zap"
)

var L *zap.Logger

func NewLogger() *zap.Logger {
	logger, _ := zap.NewProduction()
	defer logger.Sync()
	return logger
}
