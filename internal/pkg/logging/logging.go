package logging

import (
	"log"

	"github.com/sirupsen/logrus"
)

// Debugln logs a message at level Debug on the standard logger.
func Debugln(args ...interface{}) {
	if logrus.IsLevelEnabled(logrus.DebugLevel) {
		log.Println("DEBUG", args)
	}
}

// Println logs a message at level Info on the standard logger.
func Println(args ...interface{}) {
	log.Println("INFO", args)
}

// Infoln logs a message at level Warn on the standard logger.
func Infoln(args ...interface{}) {
	logrus.Infoln(args...)
}

// Warnln logs a message at level Warn on the standard logger.
func Warnln(args ...interface{}) {
	logrus.Warnln(args...)
}

// Errorln logs a message at level Error on the standard logger.
func Errorln(args ...interface{}) {
	logrus.Errorln(args...)
}

// Panicln logs a message at level Panic on the standard logger.
func Panicln(args ...interface{}) {
	logrus.Panicln(args...)
}

// Fatalln logs a message at level Fatal on the standard logger then the process will exit with status set to 1.
func Fatalln(args ...interface{}) {
	logrus.Fatalln(args...)
}
