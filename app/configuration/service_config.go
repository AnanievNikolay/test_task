package configuration

import (
	"github.com/AnanievNikolay/test_task/app/file"
)

//NewServiceConfig ...
func NewServiceConfig(_path *file.Path) *ServiceConfiguration {
	config := &ServiceConfiguration{}
	//provider := NewJSONProvider(file.NewPath(os.Args[0], "/service_config/service.json"), config)
	provider := NewJSONProvider(_path, config)
	provider.Provide()
	return config
}

//ServiceConfiguration ...
type ServiceConfiguration struct {
	ExternalAPIHost       string `json:"ExternalAPIHost"`
	MySQLConnectionString string `json:"MySQLConnectionString"`
	ServiceHost           string `json:"ServiceHost"`
	ServicePort           int    `json:"ServicePort"`
	SchedulerDuration     int    `json:"SchedulerDurationSec"`
	WSPort                int    `json:"WSPort"`
}

//ConnectionString ...
func (config *ServiceConfiguration) ConnectionString() string {
	return config.MySQLConnectionString
}

//ExternalAPI ...
func (config *ServiceConfiguration) ExternalAPI() string {
	return config.ExternalAPIHost
}

//Host ...
func (config *ServiceConfiguration) Host() string {
	return config.ServiceHost
}

//Port ..
func (config *ServiceConfiguration) Port() int {
	return config.ServicePort
}

//Duration ...
func (config *ServiceConfiguration) Duration() int {
	return config.SchedulerDuration
}

//WebsocketPort ..
func (config *ServiceConfiguration) WebsocketPort() int {
	return config.WSPort
}
