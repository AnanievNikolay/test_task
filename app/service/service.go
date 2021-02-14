package service

import (
	"log"
	"os"

	"github.com/AnanievNikolay/test_task/app/configuration"
	"github.com/AnanievNikolay/test_task/app/file"
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
	serviceConfig := configuration.NewServiceConfig(file.NewPath(os.Args[0], "/service_config/service.json"))
	settings := configuration.NewSettings(file.NewPath(os.Args[0], "/service_config/settings.yaml"))
	pool := websocket.NewPool()
	repository := repository.NewMySQLRepository(serviceConfig.ConnectionString())
	uow := unitofwork.New(repository)
	wsServer := websocket.New(serviceConfig, settings, pool.ConnectChannel(), pool.DisconnectChannel())
	client := domain.NewClient(serviceConfig.ExternalAPI(), settings.Fsyms(), settings.Tsyms())
	job := usecase.NewSchedulerJob(client, pool, uow)
	scheduler := domain.NewScheduler(serviceConfig.Duration(), job)
	return &Service{
		wsServer:   wsServer,
		uow:        uow,
		pool:       pool,
		httpServer: http.New(serviceConfig),
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

//Start ...
func (s *Service) Start() {
	defer func() {
		if r := recover(); r != nil {
			log.Println("Recovered in service start:", r)
		}
	}()
	go s.wsServer.Start()
	go s.pool.Listen()
	go s.scheduler.Start()
	s.httpServer.Start()
}
