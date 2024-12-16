package ports

import icmpdomain "github.com/rafaelpissolatto/synthetic-monitor-exporter/internal/core/domain"

// ICMPPort is the interface that wraps the DoRequest method
type ICMPPort interface {
	DoRequest(endpoint icmpdomain.ICMPEndpoint) (icmpdomain.ICMPEndpointResult, error)
}
