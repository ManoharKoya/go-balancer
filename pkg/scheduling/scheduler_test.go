package scheduling

import (
	"fmt"
	"go-balancer/internal/lbcore"
	"go-balancer/internal/lbcore/servers"
	"go-balancer/utils/stringutils"
	"go-balancer/utils/testutils"
	"testing"
	"time"
)

func TestScheduleHealthCheck(t *testing.T) {

	endpoint := "port"
	body := ""

	server := testutils.MockHttpServer(t, endpoint, body)
	defer server.Close()

	address, port := stringutils.DecodeHost(server.URL)

	seconds := 2

	var tcs = []struct {
		healthCheck    *lbcore.HealthCheck
		backendServers []*servers.BackendServer
	}{
		{
			healthCheck:    &lbcore.HealthCheck{Endpoint: endpoint, HttpBody: body, Frequency: time.Duration(seconds * 1e9)},
			backendServers: []*servers.BackendServer{{Address: address, Port: port}},
		},
		{
			healthCheck:    &lbcore.HealthCheck{Endpoint: endpoint, HttpBody: body, Schedule: fmt.Sprintf("%d * * * * *", seconds)},
			backendServers: []*servers.BackendServer{{Address: address, Port: port}},
		},
	}
	for _, tc := range tcs {
		c := ScheduleHealthCheck(tc.healthCheck, tc.backendServers)
		c.Start()
		time.Sleep(3 * time.Second)
		entries := c.Entries()
		if len(entries) != 1 {
			t.Errorf("Expected atleast one entry in health check scheduler cron.")
		} else if tc.healthCheck.Frequency != time.Duration(0) && entries[0].Next.Sub(entries[0].Prev) != time.Duration(seconds*1e9) {
			t.Errorf("Expected %d seconds of duration, but found %f seconds", seconds, entries[0].Next.Sub(entries[0].Prev).Seconds())
		}
		c.Stop()
	}
}
