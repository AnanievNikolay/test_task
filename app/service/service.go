package service

import (
	"log"

	"github.com/AnanievNikolay/test_task/app/configuration"
	"github.com/AnanievNikolay/test_task/app/servers/http"
	"github.com/AnanievNikolay/test_task/app/servers/websocket"
	"github.com/AnanievNikolay/test_task/app/servers/websocket/interfaces"
	"github.com/AnanievNikolay/test_task/datasource/repository"
	"github.com/AnanievNikolay/test_task/datasource/unitofwork"
	"github.com/AnanievNikolay/test_task/domain"
	"github.com/AnanievNikolay/test_task/usecase"
)

//New ...
func New() *Service {
	pool := websocket.NewPool()
	host := configuration.ServiceConfig().ServiceHost
	port := configuration.ServiceConfig().ServicePort
	repository := repository.NewMySQLRepository(configuration.ServiceConfig().MySQLConnectionString)
	uow := unitofwork.New(repository)
	wsServer := websocket.New(host, port, pool.ConnectChannel(), pool.DisconnectChannel())
	job := usecase.NewSchedulerJob(pool, uow)
	scheduler := domain.NewScheduler(configuration.ServiceConfig().SchedulerDuration, job)
	return &Service{
		wsServer:   wsServer,
		uow:        unitofwork.New(repository),
		pool:       pool,
		httpServer: http.New(),
		scheduler:  scheduler,
	}
}

//Service ...
type Service struct {
	wsServer   *websocket.Server
	httpServer *http.Server
	uow        unitofwork.IUnitOfWork
	pool       interfaces.IPool
	scheduler  domain.ISceduler
}

func (s *Service) Start() {
	defer func() {
		if r := recover(); r != nil {
			log.Println("Recovered in service start:", r)
		}
	}()
	go s.wsServer.Start()
	go s.pool.Listen()
	s.scheduler.Start()
}
