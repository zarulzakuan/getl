package getl

import (
	"time"

	"github.com/go-co-op/gocron"
)

type Scheduler struct {
	SchedAt *gocron.Scheduler
}

func At(cronExp string) *Scheduler {

	s := new(Scheduler)
	n := gocron.NewScheduler(time.UTC)

	s.SchedAt = n.Every(5).Seconds()

	return s
}
