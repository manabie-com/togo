package log

import (
	"ansidev.xyz/pkg/tm"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"log"
	"os"
	"strings"
	"time"
)

var (
	l *zap.Logger
	s *zap.SugaredLogger

	Debugz func(message string, fields ...zap.Field)
	Infoz  func(message string, fields ...zap.Field)
	Warnz  func(message string, fields ...zap.Field)
	Errorz func(message string, fields ...zap.Field)
	Fatalz func(message string, fields ...zap.Field)
	Panicz func(message string, fields ...zap.Field)

	Debug func(args ...interface{})
	Info  func(args ...interface{})
	Warn  func(args ...interface{})
	Error func(args ...interface{})
	Fatal func(args ...interface{})
	Panic func(args ...interface{})

	Sync func() error
)

func L() *zap.Logger {
	return l
}

// InitLogger init logger for application
// encoding: console, json
func InitLogger(encoding string) {
	level := GetLogLevel()

	config := zap.Config{
		Encoding:         encoding,
		Level:            level,
		OutputPaths:      []string{"stderr"},
		ErrorOutputPaths: []string{"stderr"},
		EncoderConfig: zapcore.EncoderConfig{
			MessageKey:   "message",
			TimeKey:      "time",
			LevelKey:     "level",
			CallerKey:    "caller",
			EncodeCaller: zapcore.ShortCallerEncoder,
			EncodeLevel:  zapcore.CapitalLevelEncoder,
			EncodeTime:   timeEncoder,
		},
	}

	l, _ = config.Build()
	s = l.Sugar()
	injectFunctions()
}

func GetLogLevel() zap.AtomicLevel {
	logLevel := os.Getenv("LOG_LEVEL")

	level, err := zap.ParseAtomicLevel(strings.ToLower(logLevel))

	if err != nil {
		log.Fatal("Could not parse log level:", err)
	}

	return level
}

func timeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format(tm.DateTimeFormat))
}

func injectFunctions() {
	Debugz = l.Debug
	Infoz = l.Info
	Warnz = l.Warn
	Errorz = l.Error
	Fatalz = l.Fatal
	Panicz = l.Panic

	Debug = s.Debug
	Info = s.Info
	Warn = s.Warn
	Error = s.Error
	Fatal = s.Fatal
	Panic = s.Panic

	Sync = l.Sync
}
