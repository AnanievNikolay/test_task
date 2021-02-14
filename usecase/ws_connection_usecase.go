package usecase

import (
	"encoding/json"
	"log"

	"github.com/AnanievNikolay/test_task/domain"
	"golang.org/x/net/websocket"
)

//NewWSConnectionUseCase ...
func NewWSConnectionUseCase(_conn *websocket.Conn, _client domain.IClient) *WSConnectionUseCase {
	return &WSConnectionUseCase{
		conn:   _conn,
		client: _client,
	}

}

//WSConnectionUseCase ...
type WSConnectionUseCase struct {
	conn   *websocket.Conn
	client domain.IClient
}

//Execute ...
func (usecase *WSConnectionUseCase) Execute() {
	response := NewPriceRequest(usecase.client).Response()
	jMess, jerr := json.Marshal(response)
	if jerr != nil {
		log.Println("[Error] Error while marshaling response. Error: ", jerr.Error())
		usecase.conn.Write([]byte("Internal server error"))
		return
	}
	usecase.conn.Write(jMess)
}
