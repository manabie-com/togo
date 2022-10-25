package gorm

import (
	"ansidev.xyz/pkg/log"
	"go.uber.org/zap"
	"gorm.io/gorm"
	gormLogger "gorm.io/gorm/logger"
	"moul.io/zapgorm2"
)

func InitGormDb(dialector gorm.Dialector) *gorm.DB {
	l := zapgorm2.New(log.L())
	l.SetAsDefault()
	zlogger := l.LogMode(getGormLogLevel())
	config := &gorm.Config{Logger: zlogger}
	gormDb, err := gorm.Open(dialector, config)

	log.FatalIf(err)

	return gormDb
}

func getGormLogLevel() gormLogger.LogLevel {
	level := log.GetLogLevel().Level()

	switch level {
	case zap.PanicLevel, zap.FatalLevel, zap.ErrorLevel:
		return gormLogger.Error
	case zap.WarnLevel:
		return gormLogger.Warn
	case zap.InfoLevel, zap.DebugLevel:
		return gormLogger.Info
	default:
		return gormLogger.Silent
	}
}
