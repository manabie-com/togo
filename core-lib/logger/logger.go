package logger

import (
	"context"
	"os"

	opentracing "github.com/opentracing/opentracing-go"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	instance *zap.Logger
)

func init() {
	config := zap.NewProductionEncoderConfig()
	config.EncodeTime = zapcore.ISO8601TimeEncoder

	core := zapcore.NewCore(zapcore.NewJSONEncoder(config), zapcore.Lock(os.Stdout), zapcore.InfoLevel)

	instance = zap.New(
		core,
		zap.AddCaller(),
		//zap.AddStacktrace(zapcore.InfoLevel),
	)
}

func For(ctx context.Context) *zap.Logger {
	// TODO: add tracing trace_id, span_id
	if span := opentracing.SpanFromContext(ctx); span != nil {
	} else {
	}
	return instance
}

func Info(msg string, fields ...zapcore.Field) {
	instance.Info(msg, fields...)
}

func Error(msg string, fields ...zapcore.Field) {
	instance.Error(msg, fields...)
}

func GetInstance() *zap.Logger {
	return instance
}
