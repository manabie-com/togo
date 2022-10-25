package log

import (
	"fmt"
	"go.uber.org/zap"
)

func WarnfIf(err error, format string, args ...interface{}) {
	if err != nil {
		l.Warn(fmt.Sprintf(format, args...), zap.Error(err))
	}
}

func WarningfIf(err error, format string, args ...interface{}) {
	WarnfIf(err, format, args...)
}

func ErrorfIf(err error, format string, args ...interface{}) {
	if err != nil {
		l.Error(fmt.Sprintf(format, args...), zap.Error(err))
	}
}

func FatalfIf(err error, format string, args ...interface{}) {
	if err != nil {
		l.Fatal(fmt.Sprintf(format, args...), zap.Error(err))
	}
}

func PanicfIf(err error, format string, args ...interface{}) {
	if err != nil {
		l.Panic(fmt.Sprintf(format, args...), zap.Error(err))
	}
}

func WarnIf(err error, args ...interface{}) {
	if err != nil {
		s.Warn(err, args)
	}
}

func WarningIf(err error, args ...interface{}) {
	WarnIf(err, args...)
}

func ErrorIf(err error, args ...interface{}) {
	if err != nil {
		s.Error(err, args)
	}
}

func FatalIf(err error, args ...interface{}) {
	if err != nil {
		s.Fatal(err, args)
	}
}

func PanicIf(err error, args ...interface{}) {
	if err != nil {
		s.Panic(err, args)
	}
}
