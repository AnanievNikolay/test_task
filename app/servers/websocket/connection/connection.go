package connection

import "golang.org/x/net/websocket"

//NewConnection ...
func NewConnection(_ip string, _conn *websocket.Conn) *Connection {
	return &Connection{
		IP:         _ip,
		Connection: _conn,
	}
}

//Connection ...
type Connection struct {
	IP         string
	Connection *websocket.Conn
}
