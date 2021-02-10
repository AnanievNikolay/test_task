package websocket

import (
	"sync"

	"golang.org/x/net/websocket"
)

//NewPool ...
func NewPool() *Pool {
	return &Pool{
		connections: make(map[string]*websocket.Conn),
		mutex:       &sync.Mutex{},
		connChan:    make(chan *Connection, 1000),
		disconnChan: make(chan string, 1000),
	}
}

//Pool ...
type Pool struct {
	connections map[string]*websocket.Conn
	mutex       *sync.Mutex
	connChan    chan *Connection
	disconnChan chan string
}

//Connect ...
func (p *Pool) Connect(_ip string, _conn *websocket.Conn) {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	p.connections[_ip] = _conn
}

//Disconnect ...
func (p *Pool) Disconnect(_ip string) {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	delete(p.connections, _ip)
}

//DisconnectChannel ...
func (p *Pool) DisconnectChannel() chan string {
	return p.disconnChan
}

//ConnectChannel ...
func (p *Pool) ConnectChannel() chan *Connection {
	return p.connChan
}

//Run ...
func (p *Pool) Listen() {
	for {
		select {
		case conn := <-p.connChan:
			{
				p.Connect(conn.ip, conn.conn)
			}
		case key := <-p.disconnChan:
			{
				p.Disconnect(key)
			}
		}
	}
}
