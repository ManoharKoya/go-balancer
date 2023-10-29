package scheduling

import (
	"fmt"
	"github.com/robfig/cron"
	"go-balancer/internal/lbcore"
	"go-balancer/internal/lbcore/servers"
	"log"
	"time"
)

func ScheduleHealthCheck(hc *lbcore.HealthCheck, s []*servers.BackendServer) *cron.Cron {
	c := cron.New()
	if hc.Frequency != time.Duration(0) {
		err := c.AddFunc(
			fmt.Sprintf("@every %s", hc.Frequency.String()),
			func() { hc.StartHealthChecks(s) })
		if err != nil {
			log.Fatalf("Error scheduling health checks %v\n", err)
		}
	} else {
		err := c.AddFunc(hc.Schedule, func() { hc.StartHealthChecks(s) })
		if err != nil {
			log.Fatalf("Error scheduling health checks %v\n", err)
		}
	}
	return c
}
