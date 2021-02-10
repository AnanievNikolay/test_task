package websocket

import (
	"fmt"
	"log"
	"net/http"

	"github.com/AnanievNikolay/test_task/app/configuration"
	"github.com/AnanievNikolay/test_task/domain"
	"github.com/AnanievNikolay/test_task/presentation/usecase"
	"golang.org/x/net/websocket"
)

//New ...
func New(_host string, _port int, _connChan chan *Connection, _disconnChan chan string) *Server {
	return &Server{
		host:        _host,
		port:        _port,
		route:       "/ws",
		connChan:    _connChan,
		disconnChan: _disconnChan,
	}
}

//Server ...
type Server struct {
	host        string
	port        int
	route       string
	connChan    chan *Connection
	disconnChan chan string
}

//Run ..
func (s *Server) Run() {
	http.Handle(s.route, websocket.Handler(s.connectionHandler))
	log.Println("Websocket server started")
	err := http.ListenAndServe(fmt.Sprintf("%v:%v", s.host, s.port), nil)
	if err != nil {
		log.Fatal("Websocket server fatal: ", err)
	}
}

func (s *Server) connectionHandler(ws *websocket.Conn) {
	connection := NewConnection(ws.RemoteAddr().String(), ws)
	log.Println("Received new connection from ", ws.RemoteAddr().String())
	s.connChan <- connection
	host := configuration.ServiceConfig().ExternalHost
	fsyms := configuration.Settings().Fsym
	tsyms := configuration.Settings().Tsym
	client := domain.NewClient(host, fsyms, tsyms)
	usecase.NewWSConnectionUseCase(ws, client).Execute()
	defer func() {
		log.Println("Client disconnected. Host : ", ws.RemoteAddr().String())
		s.disconnChan <- ws.RemoteAddr().String()
	}()
	for {
		_, err := ws.Read(make([]byte, 5))
		if err != nil {
			log.Println("[Error] Host:", ws.RemoteAddr().String(), ". Error: ", err.Error())
			return
		}
	}
}
