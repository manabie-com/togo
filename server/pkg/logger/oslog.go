package logger

import (
	"fmt"
	"log"
	"os"
)

var (
	WarningLevel         = "[WARNING]"
	ErrorLevel           = "[ERROR]"
	InfoLevel            = "[INFO]"
	DebugLevel = "[DEBUG]"
)

var logger = log.New(os.Stderr, "logger: ", log.Lshortfile)

func logWithLevel(level, msg string) {
	logger.Output(3, level + " " + msg)
}

func Warn(msg string) {
	logWithLevel(WarningLevel, msg)
}

func Info(msg string) {
	logWithLevel(InfoLevel, msg)
}

func Error(msg string) {
	logWithLevel(ErrorLevel, msg)
}

func Debug(msg string) {
	logWithLevel(ErrorLevel, msg)
}



func logfWithLevel(level, format string, v ...interface{}) {
	logger.Output(3, fmt.Sprintf(format, v...))
}

func Warnf(format string, v ...interface{}) {
	logfWithLevel(WarningLevel, format, v...)
}

func Infof(format string, v ...interface{}) {
	logfWithLevel(InfoLevel, format, v...)
}

func Debugf(format string, v ...interface{}) {
	logfWithLevel(InfoLevel, format, v...)
}

func Errorf(format string, v ...interface{}) {
	logfWithLevel(ErrorLevel, format, v...)
}