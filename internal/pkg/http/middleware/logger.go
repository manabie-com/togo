package middleware

import (
	"github.com/dinhquockhanh/togo/internal/pkg/log"
	"github.com/dinhquockhanh/togo/internal/pkg/uuid"
	"github.com/gin-gonic/gin"
)

const (
	xRequestIDKey = "X-Request-ID"
)

//SetLogger set the logger with some context inside the logger.
func SetLogger(l log.Logger) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		requestID := ctx.GetHeader(xRequestIDKey)
		req := ctx.Request
		if requestID == "" {
			requestID = uuid.New()
		}
		ctx.Header(xRequestIDKey, requestID)
		newLogger := l.WithFields(log.Fields{
			"request_id":  requestID,
			"path":        req.URL.Path,
			"remote_addr": req.RemoteAddr,
			"method":      req.Method,
		})

		ctx.Request = req.WithContext(log.NewContext(req.Context(), newLogger))
		ctx.Next()
	}
}
