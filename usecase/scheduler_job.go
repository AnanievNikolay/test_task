package usecase

import (
	"encoding/json"
	"log"
	"strings"
	"time"

	"github.com/AnanievNikolay/test_task/app/configuration"
	"github.com/AnanievNikolay/test_task/app/servers/websocket/interfaces"
	"github.com/AnanievNikolay/test_task/datasource/unitofwork"
	"github.com/AnanievNikolay/test_task/domain"
)

//NewSchedulerJob ...
func NewSchedulerJob(_pool interfaces.IPool, _uow unitofwork.IUnitOfWork) *SchedulerJob {
	return &SchedulerJob{
		pool: _pool,
		uow:  _uow,
	}
}

//SchedulerJob ...
type SchedulerJob struct {
	pool interfaces.IPool
	uow  unitofwork.IUnitOfWork
}

//Execute ...
func (job *SchedulerJob) Execute() {
	host := configuration.ServiceConfig().ExternalAPIHost
	fsyms := configuration.Settings().Fsym
	tsyms := configuration.Settings().Tsym
	response := NewPriceRequest(domain.NewClient(host, strings.Join(fsyms, ","), strings.Join(tsyms, ","))).Response()
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
