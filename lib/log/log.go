package log

import (
	"fmt"
	"os"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	logInstance *zap.SugaredLogger
)

func Get() *zap.SugaredLogger {
	return logInstance
}

func Install() {
	conf := getConfigFromEnv()
	if _, err := os.Stat(conf.Path); os.IsNotExist(err) {
		err := os.Mkdir(conf.Path, os.ModeDir)
		if err != nil {
			panic(err)
		}
	}
	file, err := os.OpenFile(conf.Path+conf.FileName, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		panic(err)
	}
	encoder := getEncoder()

	w := zapcore.NewMultiWriteSyncer(
		zapcore.AddSync(os.Stdout),
		zapcore.AddSync(file),
	)

	core := zapcore.NewCore(
		encoder,
		w,
		zapcore.DebugLevel,
	)

	logInstance = zap.New(core, zap.AddCaller()).Sugar()
	fmt.Println("====> Log Install Done!")
}

func getEncoder() zapcore.Encoder {
	return zapcore.NewConsoleEncoder(zapcore.EncoderConfig{
		MessageKey:       "msg",
		TimeKey:          "time",
		LevelKey:         "level",
		CallerKey:        "caller",
		EncodeLevel:      CustomLevelEncoder,         //Format cách hiển thị level log
		EncodeTime:       SyslogTimeEncoder,          //Format hiển thị thời điểm log
		EncodeCaller:     zapcore.ShortCallerEncoder, //Format caller
		ConsoleSeparator: " | ",
	})
}

func SyslogTimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format("2006-01-02 15:04:05"))
}

func CustomLevelEncoder(level zapcore.Level, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(level.CapitalString())
}
