package scheduler

import "github.com/robfig/cron/v3"

type Scheduler interface {
	AddJob(spec string, task func()) error
	Start()
	Stop()
}

type cronScheduler struct {
	cron *cron.Cron
}

func NewScheduler() *cronScheduler {
	return &cronScheduler{
		cron: cron.New(),
	}
}

func (s *cronScheduler) AddJob(spec string, job func()) (err error) {
	_, err = s.cron.AddFunc(spec, job)
	return
}

func (s *cronScheduler) Start() {
	s.cron.Start()
}

func (s *cronScheduler) Stop() {
	s.cron.Stop()
}
