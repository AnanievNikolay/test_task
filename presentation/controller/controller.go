package controller

import (
	"log"
	"net/http"

	_ "github.com/AnanievNikolay/test_task/docs"
	"github.com/AnanievNikolay/test_task/domain"
	"github.com/AnanievNikolay/test_task/usecase"
	"github.com/gin-gonic/gin"
)

//New ...
func New(_host string) *Controller {
	return &Controller{
		externalAPIHost: _host,
	}
}

//Controller ...
type Controller struct {
	externalAPIHost string
}

// Price godoc
// @Summary Returns cryptocurrency quotes
// @Produce json
// @Param fsyms query string true "fsyms"
// @Param tsyms query string true "tsyms"
// @Success 200 {object} model.Currency
// @Failure 404
// @Router /service/price [get]
func (c *Controller) Price(ctx *gin.Context) {
	fsyms := ctx.Query("fsyms")
	tsyms := ctx.Query("tsyms")
	response := usecase.NewPriceRequest(domain.NewClient(c.externalAPIHost, fsyms, tsyms)).Response()
	if response == nil {
		log.Println("[Controller | Price] Not found!")
		ctx.JSON(http.StatusNotFound, struct{}{})
	}
	ctx.JSON(http.StatusOK, response)
}
