package logs

import (
	"log"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var logger *Logger

func init() {
	logger = &Logger{}
	config := zap.NewProductionConfig()
	config.OutputPaths = []string{"logs.log"}
	config.DisableStacktrace = true
	config.EncoderConfig.TimeKey = "datetime"
	config.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder

	zLog, err := config.Build()
	if err != nil {
		log.Panicf("Init logger occur problem: %v", err)
	}

	logger.zLog = zLog
}

type Logger struct {
	zLog *zap.Logger
}

func WithPrefix(name string) *Logger {
	zLog := logger.zLog.Named(name)

	logger.zLog = zLog
	return logger
}

func (l *Logger) Info(content string, errorL interface{}) {
	l.zLog.Info(content,
		zap.Any("model", errorL))
}
func (l *Logger) Error(content string, errorL interface{}) {
	l.zLog.Error(content,
		zap.Any("error", errorL))
}
func (l *Logger) Panic(content string, errorL interface{}) {
	l.zLog.Panic(content,
		zap.Any("error", errorL))
}
