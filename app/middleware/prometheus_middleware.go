package middleware

import (
	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus"
	"net/http"
)

// PrometheusMiddleware /**
func PrometheusMiddleware() mux.MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			timer := prometheus.NewTimer(httpDuration.WithLabelValues(r.URL.Path))
			totalRequests.WithLabelValues(r.URL.Path).Inc()
			totalHTTPMethods.WithLabelValues(r.Method).Inc()
			next.ServeHTTP(w, r)
			timer.ObserveDuration()
		})
	}
}

func init() {
	_ = prometheus.Register(totalRequests)
	_ = prometheus.Register(totalHTTPMethods)
	_ = prometheus.Register(httpDuration)
}

var totalRequests = prometheus.NewCounterVec(
	prometheus.CounterOpts{
		Name: "http_requests_total",
		Help: "Number of incoming requests",
	},
	[]string{"path"},
)

var totalHTTPMethods = prometheus.NewCounterVec(
	prometheus.CounterOpts{
		Name: "http_methods_total",
		Help: "Number of requests per HTTP method",
	}, []string{"method"},
)

var httpDuration = prometheus.NewHistogramVec(
	prometheus.HistogramOpts{
		Name: "http_response_time_seconds",
		Help: "Duration of HTTP requests",
	},
	[]string{"path"},
)
