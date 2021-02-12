package controller

import (
	"log"
	"net/http"

	"github.com/AnanievNikolay/test_task/app/configuration"
	"github.com/AnanievNikolay/test_task/domain"
	"github.com/AnanievNikolay/test_task/usecase"
	"github.com/gin-gonic/gin"
)

//Controller ...
type Controller struct {
}

//Price ...
// @Router /service/price [get]
func (c *Controller) Price(ctx *gin.Context) {
	fsyms := ctx.Query("fsyms")
	tsyms := ctx.Query("tsyms")
	host := configuration.ServiceConfig().ExternalAPIHost
	response := usecase.NewPriceRequest(domain.NewClient(host, fsyms, tsyms)).Response()
	if response == nil {
		log.Println("[Controller | Price] Not found!")
		ctx.JSON(http.StatusNotFound, response)
	}
	ctx.JSON(http.StatusOK, response)
}
