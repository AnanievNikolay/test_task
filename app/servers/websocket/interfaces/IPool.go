package interfaces

import "github.com/AnanievNikolay/test_task/app/servers/websocket/connection"

//IPool ...
type IPool interface {
	ConnectChannel() chan *connection.Connection
	DisconnectChannel() chan string
	Listen()
	Notify([]byte)
}
