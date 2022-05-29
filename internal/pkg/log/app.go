package log

import "go.uber.org/zap"

type (
	Logger interface {
		// Infof log info with template message.
		Infof(template string, args ...any)

		// Debugf log debug with template message.
		Debugf(template string, args ...any)

		// Warnf log warning with template message.
		Warnf(template string, args ...any)

		// Errorf log error with template message.
		Errorf(template string, args ...any)

		// Panicf log with template message.
		Panicf(template string, args ...any)

		// Info log info.
		Info(args ...any)

		// Debug log debug.
		Debug(args ...any)

		// Warn log warning.
		Warn(args ...any)

		// Error log error.
		Error(args ...any)

		// Panic panic.
		Panic(args ...any)

		WithFields(fields Fields) Logger
	}

	appLogger struct {
		logger *zap.SugaredLogger
	}

	// Fields is alias of map
	Fields = map[string]any
)

// newAppLogger return new appLogger instance
func newAppLogger() *appLogger {
	return &appLogger{
		logger: newLogger(),
	}
}

// newAppLoggerWithField return a new appLogger instance with field
func newAppLoggerWithField(args Fields) *appLogger {
	return &appLogger{
		logger: newLogger().With(ToSlice(args)...),
	}
}

// Info log info
func (g *appLogger) Info(args ...any) {
	g.logger.Info(args...)
}

// Debug print debug
func (g *appLogger) Debug(args ...any) {
	g.logger.Debug(args...)
}

// Warn print warning
func (g *appLogger) Warn(args ...any) {
	g.logger.Warn(args...)
}

// Errorf print error
func (g *appLogger) Error(args ...any) {
	g.logger.Error(args...)
}

// Panic log a message, then panics.
func (g *appLogger) Panic(args ...any) {
	g.logger.Panic(args...)
}

// Infof log info with template message.
func (g *appLogger) Infof(template string, args ...any) {
	g.logger.Infof(template, args...)
}

// Debugf log debug with template message.
func (g *appLogger) Debugf(template string, args ...any) {
	g.logger.Debugf(template, args...)
}

// Warnf log warning with template message.
func (g *appLogger) Warnf(template string, args ...any) {
	g.logger.Warnf(template, args...)
}

// Errorf log error with template message.
func (g *appLogger) Errorf(template string, args ...any) {
	g.logger.Errorf(template, args...)
}

// Panicf log with template message.
func (g *appLogger) Panicf(template string, args ...any) {
	g.logger.Panicf(template, args...)
}

// WithFields return a new logger with fields
func (g *appLogger) WithFields(args Fields) Logger {

	return &appLogger{
		logger: g.logger.With(ToSlice(args)...),
	}
}

func ToSlice(f Fields) []any {
	var temp []any
	for k, v := range f {
		temp = append(temp, k, v)
	}
	return temp
}
