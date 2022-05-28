package middleware

import (
	"fmt"
	"runtime"

	"github.com/dinhquockhanh/togo/internal/pkg/http/response"
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

//Recover handle panic occurred
func Recover() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		defer func() {
			if rec := recover(); rec != nil {
				err, ok := rec.(error)
				if !ok {
					err = fmt.Errorf("%v", rec)
				}
				stack := make([]byte, 4<<10) // 4KB
				length := runtime.Stack(stack, false)

				log.WithCtx(ctx.Request.Context()).Errorf("panic recover, err: %v, stack: %s", err, stack[:length])
				response.Error(ctx, err)

			}
		}()
		ctx.Next()
	}
}
