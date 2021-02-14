package tests

import (
	"testing"
	"time"

	"github.com/AnanievNikolay/test_task/domain"
)

func TestScheduler(t *testing.T) {
	testJob := NewTestJob()
	scheduler := domain.NewScheduler(1, testJob)
	go scheduler.Start()
	time.Sleep(6 * time.Second)
	go scheduler.Stop()
	if testJob.Number != 6 {
		t.Error("Unexpected test job value. Expected: 6; Current: ", testJob.Number)
	}
}
