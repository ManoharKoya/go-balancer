package lbcore

import (
	"go-balancer/internal/lbcore/servers"
	"go-balancer/utils/stringutils"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

var healthCheck = HealthCheck{
	Endpoint:         "ping",
	HttpMethod:       "GET",
	HttpBody:         "",
	ExpectedResponse: http.Response{StatusCode: 200},
	Timeout:          time.Duration(1e9 * 2),
	Frequency:        time.Duration(1e9 * 60 * 60 * 24),
}

func TestStartHealthChecks(t *testing.T) {

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		requestUrlPath := "/" + healthCheck.Endpoint
		if req.URL.Path != requestUrlPath {
			t.Errorf("Expected to request %s, but requesting %s.", requestUrlPath, req.URL.Path)
		} else if stringutils.ReadCloserString(&req.Body) != healthCheck.HttpBody {
			t.Errorf("Expected request with %s body, but got %s.", healthCheck.HttpBody, stringutils.ReadCloserString(&req.Body))
		}
		w.WriteHeader(http.StatusOK)
	}))
	defer server.Close()

	n := 5
	var s []*servers.BackendServer
	for i := 0; i < n; i++ {
		address, port := stringutils.DecodeHost(server.URL)
		bs := servers.BackendServer{Address: address, Port: port}
		s = append(s, &bs)
	}
	healthCheck.StartHealthChecks(s)
	for i := 0; i < n; i++ {
		if s[i].IsHealthy != true {
			t.Errorf("Expected server-%d to be healthy, but found not.", i+1)
		}
	}
}
