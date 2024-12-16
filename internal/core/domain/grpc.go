package domain

import "time"

type GRPCEndpoint struct {
	Host     string
	Port     int
	Path     string
	AuthName string
	Timeout  time.Duration
	Interval time.Duration
}

type GRPCEndpointResult struct {
	Endpoint         *GRPCEndpoint
	ResponseDuration int
	ResponseBody     []byte
	ResponseError    error
	IsHealthy        bool
}

type GRPCDomain interface {
	DoRequest(endpoint GRPCEndpoint) (GRPCEndpointResult, error)
}
