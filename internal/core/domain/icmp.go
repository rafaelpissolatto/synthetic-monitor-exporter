package domain

import "time"

type ICMPEndpoint struct {
	Host     string
	Port     int
	Timeout  time.Duration
	Interval time.Duration
}

type ICMPEndpointResult struct {
	Endpoint         *ICMPEndpoint
	ResponseDuration int
	ResponseBody     []byte
	ResponseError    error
	IsHealthy        bool
}

type ICMPDomain interface {
	DoRequest(endpoint ICMPEndpoint) (ICMPEndpointResult, error)
}
