package websocket

import (
	"fmt"
	"log"
	"net/http"

	"github.com/AnanievNikolay/test_task/app/configuration"
	"github.com/AnanievNikolay/test_task/app/servers/websocket/connection"
	"github.com/AnanievNikolay/test_task/domain"
	"github.com/AnanievNikolay/test_task/usecase"
	"golang.org/x/net/websocket"
)

//New ...
func New(_config configuration.IServiceConfig, _settinggs configuration.ISettings, _connChan chan *connection.Connection, _disconnChan chan string) *Server {
	return &Server{
		config:      _config,
		settings:    _settinggs,
		connChan:    _connChan,
		disconnChan: _disconnChan,
	}
}

//Server ...
type Server struct {
	config      configuration.IServiceConfig
	settings    configuration.ISettings
	connChan    chan *connection.Connection
	disconnChan chan string
}

//Start ..
func (s *Server) Start() {
	http.Handle("/ws", websocket.Handler(s.connectionHandler))
	log.Println("Websocket server started")
	err := http.ListenAndServe(fmt.Sprintf("%v:%v", s.config.Host(), s.config.WebsocketPort()), nil)
	if err != nil {
		log.Fatal("Websocket server fatal: ", err)
	}
}

//Stop ...
func (s *Server) Stop() {
	close(s.connChan)
	close(s.disconnChan)
}

func (s *Server) connectionHandler(ws *websocket.Conn) {
	newConnection := connection.NewConnection(ws.RemoteAddr().String(), ws)
	log.Println("Received new connection from ", ws.RemoteAddr().String())
	s.connChan <- newConnection
	client := domain.NewClient(s.config.ExternalAPI(), s.settings.Fsyms(), s.settings.Tsyms())
	usecase.NewWSConnectionUseCase(ws, client).Execute()
	defer func() {
		log.Println("Client disconnected. Host : ", ws.RemoteAddr().String())
		s.disconnChan <- ws.RemoteAddr().String()
	}()
	for {
		_, err := ws.Read(make([]byte, 5))
		if err != nil {
			if err.Error() != "EOF" {
				log.Println("[Error] Host:", ws.RemoteAddr().String(), ". Error: ", err.Error())
			}
			return
		}
	}
}
