package lbcore

import (
	"bytes"
	"fmt"
	"go-balancer/internal/lbcore/servers"
	"io"
	"log"
	"net/http"
	"time"
)

// HealthCheck is the structure and details related to health checking logic for backend servers.
type HealthCheck struct {
	Endpoint         string
	HttpMethod       string
	HttpBody         string
	ExpectedResponse http.Response
	Timeout          time.Duration
	Frequency        time.Duration
	LastHealthCheck  time.Time
	Schedule         string
}

// StartHealthChecks will start the health checks on all backend servers.
func (hc *HealthCheck) StartHealthChecks(servers []*servers.BackendServer) {
	for _, s := range servers {
		if hc.PerformHealthCheck(s) {
			s.IsHealthy = true
		} else {
			s.IsHealthy = false
		}
		s.LastHealthCheck = time.Now()
	}
}

// PerformHealthCheck will perform the health check on each server
// by requests built http.Request to server and Validates http.Response.
func (hc *HealthCheck) PerformHealthCheck(server *servers.BackendServer) bool {
	client := &http.Client{Timeout: hc.Timeout}
	resp, err := client.Do(hc.httpRequest(server))
	if err != nil {
		log.Fatal("Error sending request: ", err.Error())
	}
	return resp.StatusCode == hc.ExpectedResponse.StatusCode
}

// httpRequest Builds http.Request based on hc.HttpMethod,
// http://<server.Address>/<hc.Endpoint>:<server.Port>, hc.HttpBody.
func (hc *HealthCheck) httpRequest(server *servers.BackendServer) *http.Request {
	request, err := http.NewRequest(
		hc.HttpMethod,
		fmt.Sprintf("http://%s:%d/%s", server.Address, server.Port, hc.Endpoint),
		io.NopCloser(bytes.NewReader([]byte(hc.HttpBody))))
	if err != nil {
		log.Fatal("Error in building httpRequest: ", err)
	}
	return request
}
