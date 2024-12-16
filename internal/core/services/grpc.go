package services

import "github.com/rafaelpissolatto/synthetic-monitor-exporter/internal/core/domain"

// GRPCService is a service that performs GRPC requests
type GRPCService struct {
}

// NewGRPCService creates a new GRPCService
func NewGRPCService() *GRPCService {
	return &GRPCService{}
}

// DoRequest performs an GRPC request
func (s *GRPCService) DoRequest(endpoint domain.GRPCEndpoint) (domain.GRPCEndpointResult, error) {
	return domain.GRPCEndpointResult{}, nil
}

// RunChecks runs the checks for the given endpoints (GRPC)
func (s *GRPCService) RunChecks(endpoints []domain.GRPCEndpoint) {
	// all check will be done in parallel
	for _, endpoint := range endpoints {
		go func(endpoint domain.GRPCEndpoint) {
			_, _ = s.DoRequest(endpoint)
		}(endpoint)
	}
}
