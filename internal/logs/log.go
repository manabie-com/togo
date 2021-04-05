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

func (l *Logger) Info(content, typeL string, errorL interface{}) {
	l.zLog.Info(content,
		zap.String("type", typeL),
		zap.Any("model", errorL))
}
func (l *Logger) Error(content, typeL string, errorL interface{}) {
	l.zLog.Error(content,
		zap.String("type", typeL),
		zap.Any("error", errorL))
}
func (l *Logger) Panic(content, typeL string, errorL interface{}) {
	l.zLog.Panic(content,
		zap.String("type", typeL),
		zap.Any("error", errorL))
}
