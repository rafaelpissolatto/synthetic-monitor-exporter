package ports

import httpdomain "github.com/rafaelpissolatto/synthetic-monitor-exporter/internal/core/domain"

type HTTPPort interface {
	DoRequest(endpoint httpdomain.HTTPEndpoint) (httpdomain.HTTPEndpointResult, error)
}
