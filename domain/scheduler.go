package domain

import (
	"time"
)

//NewScheduler ...
func NewScheduler(_duration int, _job IJob) ISceduler {
	return &Scheduler{
		duration:    _duration,
		job:         _job,
		stopChannel: make(chan struct{}),
	}
}

//Scheduler ...
type Scheduler struct {
	duration    int
	job         IJob
	stopChannel chan struct{}
}

//Start ...
func (s *Scheduler) Start() {
	ticker := time.NewTicker(time.Duration(s.duration) * time.Second)
	for {
		select {
		case <-ticker.C:
			{
				s.job.Execute()
			}
		case <-s.stopChannel:
			{
				ticker.Stop()
				return
			}
		}
	}
}

//Stop ...
func (s *Scheduler) Stop() {
	s.stopChannel <- struct{}{}
	close(s.stopChannel)
}
