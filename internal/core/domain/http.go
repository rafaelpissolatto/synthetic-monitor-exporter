package domain

import (
	"net/http"
	"time"
)

type HTTPEndpoint struct {
	Client   *http.Client
	Method   string // HTTP, HTTPS
	Verb     string // GET, POST, PUT, DELETE
	URL      string
	Path     string
	Port     int
	Body     []byte
	Timeout  time.Duration
	Interval time.Duration
	AuthName string // TokenManager
	Headers  map[string]string
}

type HTTPEndpointResult struct {
	Endpoint           *HTTPEndpoint
	ResponseStatusCode int
	ResponseDuration   time.Duration
	ResponseHeaders    map[string][]string
	ResponseBody       []byte
	ResponseError      error
	IsHealthy          bool
}
