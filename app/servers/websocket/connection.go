package websocket

import "golang.org/x/net/websocket"

//NewConnection ...
func NewConnection(_ip string, _conn *websocket.Conn) *Connection {
	return &Connection{
		ip:   _ip,
		conn: _conn,
	}
}

//Connection ...
type Connection struct {
	ip   string
	conn *websocket.Conn
}
