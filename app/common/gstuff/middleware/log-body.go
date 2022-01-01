package middleware

import (
	"bytes"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/manabie-com/togo/app/common/gstuff/glog"
)

type bodyDumpResponseWriter struct {
	io.Writer
	http.ResponseWriter
}

func (w *bodyDumpResponseWriter) Write(b []byte) (int, error) {
	return w.Writer.Write(b)
}

type MyResponseBuffer struct {
	B *bytes.Buffer
}

func (b *MyResponseBuffer) Write(p []byte) (n int, err error) {
	if len(p)+b.B.Len() < 5000 {
		b.B.Write(p)
	}
	return len(p), nil
}

func NewResponseBuffer() *MyResponseBuffer {
	return &MyResponseBuffer{B: new(bytes.Buffer)}
}

// LogBody ..
func LogBody(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) (err error) {

		req := c.Request()
		res := c.Response()
		start := time.Now()

		// set requestID
		c.Set("reqID", res.Header().Get(echo.HeaderXRequestID))

		// create data send elastic
		data := map[string]interface{}{
			"method":     req.Method,
			"headers":    req.Header,
			"status":     res.Status,
			"remote-ip":  c.RealIP(),
			"user-agent": req.UserAgent(),
			"request-id": res.Header().Get(echo.HeaderXRequestID),
			"uri":        req.RequestURI,
		}

		// get body-request
		{
			bodyRequest := []byte{}
			contentType := req.Header.Get("Content-Type")
			if strings.Contains(contentType, "application/json") {
				if req.Body != nil { // Read
					bodyRequest, _ = ioutil.ReadAll(io.LimitReader(req.Body, 50*1024)) // limit read 50kb body request
				}

				// reset
				req.Body = ioutil.NopCloser(bytes.NewBuffer(bodyRequest))

				bodyReqSize := len(bodyRequest)

				bodyRequestLimit := string(bodyRequest)
				if len(bodyRequestLimit) > 5000 {
					bodyRequestLimit = bodyRequestLimit[:5000]
				}

				data["body-request"] = bodyRequestLimit
				data["body-request-size"] = bodyReqSize

				if bodyReqSize >= 50*1024 {
					// update status
					data["status"] = http.StatusRequestEntityTooLarge
					data["body-response"] = "Body request too large"

					// send to elastic
					glog.Send(data)

					return echo.NewHTTPError(http.StatusRequestEntityTooLarge, "Body request too large")
				}
			}
		}

		// get body-response
		{
			bodyResponse := NewResponseBuffer()
			mw := io.MultiWriter(res.Writer, bodyResponse)
			writer := &bodyDumpResponseWriter{Writer: mw, ResponseWriter: res.Writer}
			res.Writer = writer

			if err := next(c); err != nil {
				c.Error(err)
			}

			bodyResSize := len(bodyResponse.B.String())

			bodyResponseLimit := ""
			contentTypeResponse := c.Response().Header().Get("Content-Type")
			if strings.Contains(contentTypeResponse, "application/json") {
				bodyResponseLimit = bodyResponse.B.String()
				if len(bodyResponseLimit) > 5000 {
					bodyResponseLimit = bodyResponseLimit[:5000]
				}
			}

			data["body-response"] = bodyResponseLimit
			data["body-response-size"] = bodyResSize
			data["status"] = res.Status
		}

		stop := time.Now()

		data["latency-human"] = stop.Sub(start).String()
		data["latency-micro"] = stop.Sub(start).Microseconds()

		// test send data to fluentd
		glog.Send(data)

		return nil
	}
}
