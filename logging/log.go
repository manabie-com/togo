package logging

import (
	"log"
	"strings"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var Logger *zap.SugaredLogger

//Init init logger with default
func init() { //*zap.SugaredLogger {
	cfg := buildConfig(getLevel("default"))
	tempLog, err := cfg.Build()
	if err != nil {
		log.Fatal("Can't init logger.")
	}
	Logger = tempLog.Sugar()
	// return tempLog.Sugar()
}

//InitWithOption set level log, service name
func InitWithOption(level, service string) (*zap.SugaredLogger, error) {
	// var err error
	cfg := buildConfig(getLevel(level))
	tempLog, err := cfg.Build()
	if err != nil {
		return nil, err
	}
	zap.ReplaceGlobals(Logger.Desugar())
	return tempLog.Sugar().With("service", service), nil
}

func buildConfig(logLevel zap.AtomicLevel) zap.Config {
	return zap.Config{
		Encoding: "json",
		Level:    logLevel,
		// set stdout & stderr seprate for easy tracking log in future
		OutputPaths:      []string{"stdout"},
		ErrorOutputPaths: []string{"stderr"},
		EncoderConfig: zapcore.EncoderConfig{
			TimeKey:    "time",
			EncodeTime: zapcore.ISO8601TimeEncoder,
			MessageKey: "message",

			LevelKey:    "level",
			EncodeLevel: zapcore.LowercaseLevelEncoder,

			CallerKey:    "caller",
			EncodeCaller: zapcore.ShortCallerEncoder,
		},
	}
}

func getLevel(level string) zap.AtomicLevel {
	level = strings.ToLower(level)
	var logLevel zap.AtomicLevel
	switch level {
	case "debug":
		logLevel = zap.NewAtomicLevelAt(zap.DebugLevel)
	case "info":
		logLevel = zap.NewAtomicLevelAt(zap.InfoLevel)
	case "warn":
		logLevel = zap.NewAtomicLevelAt(zap.WarnLevel)
	case "error":
		logLevel = zap.NewAtomicLevelAt(zap.ErrorLevel)
	default:
		logLevel = zap.NewAtomicLevelAt(zap.InfoLevel)
	}
	return logLevel
}
