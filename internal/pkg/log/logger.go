package log

import (
	"context"
	"sync"

	"github.com/gin-gonic/gin"
)

var (
	mainLogger *appLogger
	once       sync.Once
)

const (
	Key = "logger"
)

func Root() *appLogger {
	once.Do(func() {
		mainLogger = newAppLogger()
	})
	return mainLogger

}

// Init the mainLogger logger with fields
func Init(fields Fields) {
	mainLogger = newAppLoggerWithField(fields)
}

// NewContext return a new logger context
func NewContext(ctx context.Context, logger Logger) context.Context {
	if logger == nil {
		logger = Root()
	}
	return context.WithValue(ctx, Key, logger)
}

// FromCtx get logger form context
func FromCtx(ctx context.Context) Logger {
	if ctx == nil {
		return Root()
	}
	if logger, ok := ctx.Value(Key).(Logger); ok {
		return logger
	}
	return Root()
}

// Info log info with template.
func Info(v ...any) {
	Root().Info(v)
}

// Infof log info with template.
func Infof(template string, v ...any) {
	Root().Infof(template, v...)
}

// Debugf log debug with template.
func Debugf(template string, v ...any) {
	Root().Debugf(template, v...)
}

// Warnf log warning with template.
func Warnf(template string, v ...any) {
	Root().Warnf(template, v...)
}

// Errorf log error with template.
func Errorf(template string, v ...any) {
	Root().Errorf(template, v...)
}

// Panicf panic with template.
func Panicf(template string, v ...any) {
	Root().Panicf(template, v...)
}

// With return a new logger entry with fields
func With(fields Fields) Logger {
	return Root().WithFields(fields)
}

// WithCtx return a logger from the given context
func WithCtx(ctx context.Context) Logger {
	if gc, ok := ctx.(*gin.Context); ok {
		ctx = gc.Request.Context()
	}
	return FromCtx(ctx)
}
