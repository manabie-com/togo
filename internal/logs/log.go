package logs

import (
	"log"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var logger *zap.Logger

func init() {
	config := zap.NewProductionConfig()
	config.OutputPaths = []string{"logs.log"}
	config.EncoderConfig.TimeKey = "datetime"
	config.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	var err error
	logger, err = config.Build()
	if err != nil {
		log.Panicf("Init logger occur problem: %v", err)
	}
}

func WithPrefix(name string) *zap.Logger {
	l := logger.Named(name)
	return l
}
