package config

import (
	"fmt"
	"os"

	"github.com/rafaelpissolatto/synthetic-monitor-exporter/internal/core/domain"
	"gopkg.in/yaml.v2"
)

// Config is the configuration for the synthetic-monitor-exporter
type Config struct {
	// ServerConfig is the configuration for the server
	ServerConfig ServerConfig `yaml:"server"`

	// HTTPConfig is the configuration for the HTTP service
	HTTPConfig HTTPConfig `yaml:"http"`

	// ICMPConfig is the configuration for the ICMP service
	ICMPConfig ICMPConfig `yaml:"icmp"`

	// GRPCConfig is the configuration for the gRPC service
	GRPCConfig GRPCConfig `yaml:"grpc"`
}

// HTTPConfig is the configuration for the HTTP service
type HTTPConfig struct {
	// Endpoints is a list of HTTP endpoints to monitor
	Endpoints []domain.HTTPEndpoint `yaml:"http_endpoints,omitempty"`
}

// ICMPConfig is the configuration for the ICMP service
type ICMPConfig struct {
	// Endpoints is a list of ICMP endpoints to monitor
	Endpoints []domain.ICMPEndpoint `yaml:"icmp_endpoints,omitempty"`
}

// GRPCConfig is the configuration for the gRPC service
type GRPCConfig struct {
	// Endpoints is a list of gRPC endpoints to monitor
	Endpoints []domain.GRPCEndpoint `yaml:"grpc_endpoints,omitempty"`
}

type ServerConfig struct {
	// Port is the port the server will listen on
	Port int `yaml:"port"`
	// LogLevel is the log level for the server
	LogLevel string `yaml:"log_level"`
}

// NewConfig creates a new Config
func NewConfig(filePath string) (*Config, error) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to read file: %w", err)
	}

	var config Config
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal config: %w", err)
	}

	return &config, nil
}
