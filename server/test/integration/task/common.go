package task

import (
	"net/http"
	"net/url"
	"time"
)

const (
	TrueToken = ""
	ErrorToken = ""
)

var TaskURL *url.URL
var httpClient *http.Client

func init() {
	TaskURL, _ = url.Parse("http://localhost:8080/api/tasks")
	httpClient = &http.Client{
		Timeout: 10 * time.Second,
	}
}
