package websocket

import (
	"log"
	"sync"

	"github.com/AnanievNikolay/test_task/app/servers/websocket/connection"
	"golang.org/x/net/websocket"
)

//NewPool ...
func NewPool() *Pool {
	return &Pool{
		connections: make(map[string]*websocket.Conn),
		mutex:       &sync.Mutex{},
		connChan:    make(chan *connection.Connection, 1000),
		disconnChan: make(chan string, 1000),
	}
}

//Pool ...
type Pool struct {
	connections map[string]*websocket.Conn
	mutex       *sync.Mutex
	connChan    chan *connection.Connection
	disconnChan chan string
}

//Notify ...
func (p *Pool) Notify(msg []byte) {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	for _, conn := range p.connections {
		_, err := conn.Write(msg)
		if err != nil {
			log.Printf("[Error] Error while send message to WS. Host:%v. Error:%v", conn.RemoteAddr().String(), err.Error())
		}
	}
}

func (p *Pool) connect(_ip string, _conn *websocket.Conn) {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	p.connections[_ip] = _conn
}

func (p *Pool) disconnect(_ip string) {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	delete(p.connections, _ip)
}

//DisconnectChannel ...
func (p *Pool) DisconnectChannel() chan string {
	return p.disconnChan
}

//ConnectChannel ...
func (p *Pool) ConnectChannel() chan *connection.Connection {
	return p.connChan
}

//Run ...
func (p *Pool) Listen() {
	for {
		select {
		case conn := <-p.connChan:
			{
				p.connect(conn.IP, conn.Connection)
			}
		case key := <-p.disconnChan:
			{
				p.disconnect(key)
			}
		}
	}
}
