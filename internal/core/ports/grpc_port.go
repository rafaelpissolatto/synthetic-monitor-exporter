package ports

import grpcdomain "github.com/rafaelpissolatto/synthetic-monitor-exporter/internal/core/domain"

type GRPCPort interface {
	DoRequest(endpoint grpcdomain.GRPCEndpoint) (grpcdomain.GRPCEndpointResult, error)
}
