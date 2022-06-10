package server

import "togo/pkg/logger"

type ServerContext interface {
	GetService(prefix string) interface{}
	GetLogger() logger.Loggers
}
