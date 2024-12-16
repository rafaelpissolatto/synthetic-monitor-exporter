package services

import "github.com/rafaelpissolatto/synthetic-monitor-exporter/internal/core/domain"

// ICMPService is a service that performs ICMP requests
type ICMPService struct {
}

// NewICMPService creates a new ICMPService
func NewICMPService() *ICMPService {
	return &ICMPService{}
}

// DoRequest performs an ICMP request
func (s *ICMPService) DoRequest(endpoint domain.ICMPEndpoint) (domain.ICMPEndpointResult, error) {
	return domain.ICMPEndpointResult{}, nil
}

// RunChecks runs the checks for the given endpoints (ICMP)
func (s *ICMPService) RunChecks(endpoints []domain.ICMPEndpoint) {
	// all check will be done in parallel
	for _, endpoint := range endpoints {
		go func(endpoint domain.ICMPEndpoint) {
			_, _ = s.DoRequest(endpoint)
		}(endpoint)
	}
}
