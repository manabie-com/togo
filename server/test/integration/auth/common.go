package auth

import (
	"net/http"
	"net/url"
	"time"
)

var AuthURL *url.URL
var httpClient *http.Client

func init() {
	AuthURL, _ = url.Parse("http://localhost:8080/api/auth")
	httpClient = &http.Client{
		Timeout: 10 * time.Second,
	}
}
