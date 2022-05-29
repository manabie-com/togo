package log

import (
	"os"

	"github.com/dinhquockhanh/togo/internal/pkg/config"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

func newLogger() *zap.SugaredLogger {
	encoder := getEncoder()
	ws := getWriteSyncer(config.All.Log.File)
	l := zapcore.DebugLevel
	switch config.All.Log.Level {
	case "info":
		l = zapcore.InfoLevel
	case "warn":
		l = zapcore.WarnLevel
	case "error":
		l = zapcore.ErrorLevel
	}

	core := zapcore.NewCore(encoder, ws, l)
	logger := zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1))
	return logger.Sugar()
}

func getEncoder() zapcore.Encoder {
	c := zap.NewProductionEncoderConfig()
	if config.All.Env == "dev" {
		c = zap.NewDevelopmentEncoderConfig()
	}

	if config.All.Log.Encoder == "console" {
		return zapcore.NewConsoleEncoder(c)
	}
	return zapcore.NewJSONEncoder(c)

}

//getWriteSyncer return syncer stdout and file if config have file name
func getWriteSyncer(filename string) zapcore.WriteSyncer {
	ws := make([]zapcore.WriteSyncer, 0)
	if filename != "" {
		var fileLogger = &lumberjack.Logger{
			Filename:   filename,
			MaxSize:    1,  // MB
			MaxBackups: 3,  // number of backups
			MaxAge:     28, //days
			LocalTime:  true,
			Compress:   false, // disabled by default
		}
		ws = append(ws, zapcore.AddSync(fileLogger))
	}
	ws = append(ws, zapcore.AddSync(os.Stdout))
	return zapcore.NewMultiWriteSyncer(ws...)
}
