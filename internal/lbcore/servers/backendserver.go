package servers

import (
	"sync"
	"time"
)

// BackendServer structure containing details for backend server.
type BackendServer struct {
	Address         string
	Port            int
	IsHealthy       bool
	LastHealthCheck time.Time
	Mutex           sync.Mutex
	MessageChan     chan string
}

// New backend server creation, returning *BackendServer
func New(address string, port int) *BackendServer {
	return &BackendServer{
		Address:   address,
		Port:      port,
		IsHealthy: false,
	}
}
