package services

import (
	"log"

	"github.com/rafaelpissolatto/synthetic-monitor-exporter/internal/config"
	"github.com/rafaelpissolatto/synthetic-monitor-exporter/internal/core/domain"
)

// HTTPService is a service that performs HTTP requests
type HTTPService struct {
	LoggingEnabled bool
}

// NewHTTPService creates a new HTTPService
func NewHTTPService() *HTTPService {
	return &HTTPService{}
}

func (s *HTTPService) DoRequest(endpoint domain.HTTPEndpoint) (domain.HTTPEndpointResult, error) {
	log.Printf("Performing HTTP request to %s\n", endpoint.URL)
	return domain.HTTPEndpointResult{}, nil
}

// RunChecks performs HTTP checks on the provided endpoints
func (s *HTTPService) RunChecks(endpoints []domain.HTTPEndpoint) {
	// all check will be done in parallel
	for _, endpoint := range endpoints {
		log.Printf("Checking HTTP endpoint %s\n", endpoint.URL)
		localEndpoint := endpoint
		go func(endpoint domain.HTTPEndpoint) {
			_, _ = s.DoRequest(endpoint)
		}(localEndpoint)
	}
}

// StartCheckerHTTPServer starts the Checker HTTP server
// config parameter contains the configuration for the HTTP server, including endpoints to check
func (s *HTTPService) StartCheckerHTTPServer(config config.Config) {
	log.Println("Starting HTTP checker server")
	endpoints := config.HTTPConfig.Endpoints
	s.RunChecks(endpoints)
}
