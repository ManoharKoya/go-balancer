package algorithms

import "go-balancer/internal/lbcore/servers"

type RoundRobinLoadBalancer struct {
	servers []servers.BackendServer
}

func New(servers []servers.BackendServer) *RoundRobinLoadBalancer {
	return &RoundRobinLoadBalancer{servers: servers}
}

func (lb *RoundRobinLoadBalancer) SelectServer() *servers.BackendServer {
	return &lb.servers[0]
}
