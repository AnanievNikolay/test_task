package usecase

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"

	"github.com/AnanievNikolay/test_task/app/configuration"
	"github.com/AnanievNikolay/test_task/domain"
	"github.com/gin-gonic/gin"
)

//NewGetPriceUseCase ...
func NewGetPriceUseCase(_ctx *gin.Context, _client domain.IClient) *GetPriceUseCase {
	return &GetPriceUseCase{
		ctx:    _ctx,
		client: _client,
	}

}

//GetPriceUseCase ...
type GetPriceUseCase struct {
	ctx    *gin.Context
	client domain.IClient
}

//Execute ...
func (usecase *GetPriceUseCase) Execute() {
	host := configuration.ServiceConfig().ExternalAPIHost
	fsyms := configuration.Settings().Fsym
	tsyms := configuration.Settings().Tsym
	response := NewPriceRequest(domain.NewClient(host, strings.Join(fsyms, ","), strings.Join(tsyms, ","))).Response()
	jMess, jerr := json.Marshal(response)
	if jerr != nil {
		log.Println("[Error] Error while marshaling response. Error: ", jerr.Error())
		usecase.ctx.JSON(http.StatusInternalServerError, "Internal server error")
		return
	}
	usecase.ctx.JSON(http.StatusOK, string(jMess))
}
