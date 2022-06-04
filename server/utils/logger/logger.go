package logger

import (
	"os"

	"github.com/rs/zerolog"
)

func NewLogger(level string) zerolog.Logger {
	// Set default logging level to `info`
	logLevel := zerolog.InfoLevel

	// Set to different logging level if `level` == "debug"
	// Otherwise set logging level to `info`
	switch level {
	case "debug":
		logLevel = zerolog.DebugLevel
	default:
		logLevel = zerolog.InfoLevel
	}

	zerolog.SetGlobalLevel(logLevel)

	// Output logs to Stdout
	// Reference: https://12factor.net/logs
	logger := zerolog.New(os.Stdout).With().Timestamp().Logger()

	return logger
}
