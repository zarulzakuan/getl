package getl

import (
	"log"
	"time"

	"github.com/go-co-op/gocron"
)

// Scheduler is a timer wrapper so we can pass it to the Source node
type Scheduler struct {
	SchedAt *gocron.Scheduler
}

// RunAt - creates a scheduler. Requires the timer in either cron expression in string
// or seconds in integer, timezone, and boolean; if true, set the job to not start immediately
// but rather wait until the first scheduled interval
func RunAt(timer any, tz *time.Location, wait bool) *Scheduler {

	s := new(Scheduler)
	n := gocron.NewScheduler(tz)

	switch timer := timer.(type) {
	case int, int32, int64:
		s.SchedAt = n.Every(timer).Seconds()
	case string:
		s.SchedAt = n.Cron(timer)
	default:
		log.Fatalln("Wrong type")
	}

	if wait {
		s.SchedAt.WaitForSchedule()
	}

	return s
}

// RunNow - Run the pipeline immediately with no scheduling
func RunNow() *Scheduler {
	return nil
}
