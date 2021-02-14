package usecase

import (
	"encoding/json"
	"log"
	"time"

	"github.com/AnanievNikolay/test_task/app/servers/websocket/interfaces"
	"github.com/AnanievNikolay/test_task/datasource/unitofwork"
	"github.com/AnanievNikolay/test_task/domain"
)

//NewSchedulerJob ...
func NewSchedulerJob(_client domain.IClient, _pool interfaces.IPool, _uow unitofwork.IUnitOfWork) *SchedulerJob {
	return &SchedulerJob{
		pool:   _pool,
		uow:    _uow,
		client: _client,
	}
}

//SchedulerJob ...
type SchedulerJob struct {
	pool   interfaces.IPool
	uow    unitofwork.IUnitOfWork
	client domain.IClient
}

//Execute ...
func (job *SchedulerJob) Execute() {
	response := NewPriceRequest(job.client).Response()
	job.uow.NewDirty(response)
	jMess, jerr := json.Marshal(response)
	if jerr != nil {
		log.Println("[Error] Error while marshaling response. Error: ", jerr.Error())
	} else {
		job.pool.Notify(jMess)
	}
	job.uow.Commit()
	log.Println("[SchedulerJob | Execute] Scheduler job executed at : ", time.Now().UTC())
}
