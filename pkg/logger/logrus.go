package pkg_logrus

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"manabie-com/togo/util"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"sync"
	"time"
)

var Lgrus = &customLogrus{}

func InitLogrus() {
	createLogrus()
}

func GetFileName() string {
	return fmt.Sprintf("%s-interview.log", time.Now().Format(util.DefaultTimeFormat))
}

func createLogrus() {
	logger := logrus.New()

	//Check if logs folder exists, create if not
	logsPath := filepath.Join(".", "logs")
	err := os.MkdirAll(logsPath, 0777)
	if err != nil {
		fmt.Println("logs path Err:", err)
	}

	//Currently only log out to one file
	fileName := GetFileName()
	filePath := filepath.Join("logs", fileName)
	file, err := os.OpenFile(filePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err == nil {
		logger.Out = file
	} else {
		logger.Infof("Failed to log to file with err : %v, using default stderr\n", err)
	}

	//Show line number and function name
	logger.SetReportCaller(false)

	Lgrus.Instance = logger
	Lgrus.FileName = fileName
}

// Logger collects logging information at several levels
type Logger interface {
	Debug(v ...interface{})
	Info(v ...interface{})
	Warning(v ...interface{})
	Error(v ...interface{})
	Critical(v ...interface{})
	Fatal(v ...interface{})
	Close()
	Rotate()
	CallFileInfo() string
}

type customLogrus struct {
	FileName string
	Instance *logrus.Logger
	mu       sync.Mutex
}

// Debug logs a message using DEBUG as log level.
func (l *customLogrus) Debug(v ...interface{}) {
	l.mu.Lock()
	l.Rotate()
	l.Instance.WithField("file", l.CallFileInfo()).Debugln(v)
	l.mu.Unlock()
}

// Info logs a message using INFO as log level.
func (l *customLogrus) Info(v ...interface{}) {
	l.mu.Lock()
	l.Rotate()
	l.Instance.WithField("file", l.CallFileInfo()).Info(v)
	l.mu.Unlock()
}

// Warning logs a message using WARNING as log level.
func (l *customLogrus) Warning(v ...interface{}) {
	l.mu.Lock()
	l.Rotate()
	l.Instance.WithField("file", l.CallFileInfo()).Warningln(v)
	l.mu.Unlock()
}

// Error logs a message using ERROR as log level.
func (l *customLogrus) Error(v ...interface{}) {
	l.mu.Lock()
	l.Rotate()
	l.Instance.WithField("file", l.CallFileInfo()).Errorln(v)
	l.mu.Unlock()
}

// Critical logs a message using CRITICAL as log level.
func (l *customLogrus) Critical(v ...interface{}) {
	l.mu.Lock()
	l.Rotate()
	l.Instance.WithField("file", l.CallFileInfo()).WithField("trace", "CRITICAL").Errorln(v)
	l.mu.Unlock()
}

// Fatal is equivalent to l.Critical(fmt.Sprint()) followed by a call to os.Exit(1).
func (l *customLogrus) Fatal(v ...interface{}) {
	l.mu.Lock()
	l.Rotate()
	l.Instance.WithField("file", l.CallFileInfo()).Fatalln(v)
	l.mu.Unlock()
}

func (l *customLogrus) Close() {
	l.Instance = nil
	l.FileName = ""
}

func (l *customLogrus) Rotate() {
	if GetFileName() != l.FileName {
		createLogrus()
	}
}

func (l *customLogrus) CallFileInfo() string {
	_, file, line, ok := runtime.Caller(2)
	if !ok {
		return ""
	} else {
		index := strings.Index(file, "interview")
		if index != -1 {
			return fmt.Sprintf(`%s:%d`, file[index:], line)
		}

		return fmt.Sprintf(`%s:%d`, file, line)
	}
}