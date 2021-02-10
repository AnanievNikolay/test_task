package usecase

import (
	"encoding/json"
	"log"

	"github.com/AnanievNikolay/test_task/app/configuration"
	"github.com/AnanievNikolay/test_task/domain"
	"golang.org/x/net/websocket"
)

//NewWSConnectionUseCase ...
func NewWSConnectionUseCase(_conn *websocket.Conn, _client domain.IClient) *WSConnectionUseCase {
	return &WSConnectionUseCase{
		conn: _conn,
	}

}

//WSConnectionUseCase ...
type WSConnectionUseCase struct {
	conn   *websocket.Conn
	client domain.IClient
}

//Execute ...
func (usecase *WSConnectionUseCase) Execute() {
	host := configuration.ServiceConfig().ExternalHost
	fsyms := configuration.Settings().Fsym
	tsyms := configuration.Settings().Tsym
	response := NewPriceRequest(domain.NewClient(host, fsyms, tsyms)).Response()
	jMess, jerr := json.Marshal(response)
	if jerr != nil {
		log.Println("[Error] Error while marshaling response. Error: ", jerr.Error())
		usecase.conn.Write([]byte("Internal server error"))
		return
	}
	usecase.conn.Write(jMess)
}
