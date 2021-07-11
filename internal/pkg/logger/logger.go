package logger

import (
	"context"
	"log"
)

const (
	INFO_LEVEL  = "info"
	DEBUG_LEVEL = "debug"
	WARN_LEVEL  = "warn"

	MBLoggerConText = "mbLogger"
)

type (
	Logger interface {
		WithFields(fields map[string]interface{}) Logger
		WithPrefix(prefix string) Logger

		Debugf(format string, args ...interface{})
		Infof(format string, args ...interface{})
		Printf(format string, args ...interface{})
		Warnf(format string, args ...interface{})
		Errorf(format string, args ...interface{})
		Panicf(format string, args ...interface{})

		Debug(args ...interface{})
		Info(args ...interface{})
		Print(args ...interface{})
		Warn(args ...interface{})
		Error(args ...interface{})
		Panic(args ...interface{})

		Debugln(args ...interface{})
		Infoln(args ...interface{})
		Println(args ...interface{})
		Warnln(args ...interface{})
		Errorln(args ...interface{})
		Panicln(args ...interface{})

		MBDebugf(ctx context.Context, format string, args ...interface{})
		MBInfof(ctx context.Context, format string, args ...interface{})
		MBPrintf(ctx context.Context, format string, args ...interface{})
		MBWarnf(ctx context.Context, format string, args ...interface{})
		MBErrorf(ctx context.Context, format string, args ...interface{})
		MBPanicf(ctx context.Context, format string, args ...interface{})

		MBDebug(ctx context.Context, args ...interface{})
		MBInfo(ctx context.Context, args ...interface{})
		MBPrint(ctx context.Context, args ...interface{})
		MBWarn(ctx context.Context, args ...interface{})
		MBError(ctx context.Context, args ...interface{})
		MBPanic(ctx context.Context, args ...interface{})

		MBDebugln(ctx context.Context, args ...interface{})
		MBInfoln(ctx context.Context, args ...interface{})
		MBPrintln(ctx context.Context, args ...interface{})
		MBWarnln(ctx context.Context, args ...interface{})
		MBErrorln(ctx context.Context, args ...interface{})
		MBPanicln(ctx context.Context, args ...interface{})
	}
)

var std Logger

func init() {
	var err error = nil
	std, err = newLogrusLogger(INFO_LEVEL)
	if err != nil {
		log.Panic(err)
	}
}

func WithFields(fields map[string]interface{}) Logger {
	return std.WithFields(fields)
}

func WithPrefix(prefix string) Logger {
	return std.WithPrefix(prefix)
}

func WithMetricType(metricType string) Logger {
	return std.WithFields(map[string]interface{}{
		"metric_type": metricType,
	})
}

func Debugf(format string, args ...interface{}) {
	std.Debugf(format, args...)
}
func Infof(format string, args ...interface{}) {
	std.Infof(format, args...)
}
func Printf(format string, args ...interface{}) {
	std.Printf(format, args...)
}
func Warnf(format string, args ...interface{}) {
	std.Warnf(format, args...)
}
func Errorf(format string, args ...interface{}) {
	std.Errorf(format, args...)
}
func Panicf(format string, args ...interface{}) {
	std.Panicf(format, args...)
}

func Debug(args ...interface{}) {
	std.Debug(args...)
}
func Info(args ...interface{}) {
	std.Info(args...)
}
func Print(args ...interface{}) {
	std.Print(args...)
}
func Warn(args ...interface{}) {
	std.Warn(args...)
}
func Error(args ...interface{}) {
	std.Error(args...)
}
func Panic(args ...interface{}) {
	std.Panic(args...)
}

func Debugln(args ...interface{}) {
	std.Debugln(args...)
}
func Infoln(args ...interface{}) {
	std.Infoln(args...)
}
func Println(args ...interface{}) {
	std.Println(args...)
}
func Warnln(args ...interface{}) {
	std.Warnln(args...)
}
func Errorln(args ...interface{}) {
	std.Errorln(args...)
}
func Panicln(args ...interface{}) {
	std.Panicln(args...)
}

func MBDebugf(ctx context.Context, format string, args ...interface{}) {
	std.MBDebugf(ctx, format, args...)
}
func MBInfof(ctx context.Context, format string, args ...interface{}) {
	std.MBInfof(ctx, format, args...)
}
func MBPrintf(ctx context.Context, format string, args ...interface{}) {
	std.MBPrintf(ctx, format, args...)
}
func MBWarnf(ctx context.Context, format string, args ...interface{}) {
	std.MBWarnf(ctx, format, args...)
}
func MBErrorf(ctx context.Context, format string, args ...interface{}) {
	std.MBErrorf(ctx, format, args...)
}
func MBPanicf(ctx context.Context, format string, args ...interface{}) {
	std.MBPanicf(ctx, format, args...)
}

func MBDebug(ctx context.Context, args ...interface{}) {
	std.MBDebug(ctx, args...)
}
func MBInfo(ctx context.Context, args ...interface{}) {
	std.MBInfo(ctx, args...)
}
func MBPrint(ctx context.Context, args ...interface{}) {
	std.MBPrint(ctx, args...)
}
func MBWarn(ctx context.Context, args ...interface{}) {
	std.MBWarn(ctx, args...)
}
func MBError(ctx context.Context, args ...interface{}) {
	std.MBError(ctx, args...)
}
func MBPanic(ctx context.Context, args ...interface{}) {
	std.MBPanic(ctx, args...)
}

func MBDebugln(ctx context.Context, args ...interface{}) {
	std.MBDebugln(ctx, args...)
}
func MBInfoln(ctx context.Context, args ...interface{}) {
	std.MBInfoln(ctx, args...)
}
func MBPrintln(ctx context.Context, args ...interface{}) {
	std.MBPrintln(ctx, args...)
}
func MBWarnln(ctx context.Context, args ...interface{}) {
	std.MBWarnln(ctx, args...)
}
func MBErrorln(ctx context.Context, args ...interface{}) {
	std.MBErrorln(ctx, args...)
}
func MBPanicln(ctx context.Context, args ...interface{}) {
	std.MBPanicln(ctx, args...)
}
