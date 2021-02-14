package http

import (
	"fmt"
	"log"

	"github.com/AnanievNikolay/test_task/app/configuration"
	"github.com/AnanievNikolay/test_task/presentation/controller"
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

//New ..
func New(_config configuration.IServiceConfig) *Server {
	ginEngine := gin.Default()
	c := controller.New(_config.ExternalAPI())
	ginEngine.GET("/service/price", c.Price)
	ginEngine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	return &Server{
		config: _config,
		server: ginEngine,
	}
}

//Server ...
type Server struct {
	server *gin.Engine
	config configuration.IServiceConfig
}

//Start ...
func (s *Server) Start() {
	log.Println("HTTP server started")
	s.server.Run(fmt.Sprintf("%v:%v", s.config.Host(), s.config.Port()))
}
