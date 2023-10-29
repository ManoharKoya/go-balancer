package testutils

import (
	"go-balancer/utils/stringutils"
	"net/http"
	"net/http/httptest"
	"testing"
)

func MockHttpServer(t *testing.T, endpoint string, body string) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		requestUrlPath := "/" + endpoint
		if req.URL.Path != requestUrlPath {
			t.Errorf("Expected to request %s, but requesting %s.", requestUrlPath, req.URL.Path)
		} else if stringutils.ReadCloserString(&req.Body) != body {
			t.Errorf("Expected request with %s body, but got %s.", body, stringutils.ReadCloserString(&req.Body))
		}
		w.WriteHeader(http.StatusOK)
	}))
}
