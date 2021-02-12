package configuration

import (
	"os"

	"github.com/AnanievNikolay/test_task/app/file"
)

var config *ServiceConfiguration

//ServiceConfig ...
func ServiceConfig() *ServiceConfiguration {
	if config == nil {
		config = &ServiceConfiguration{}
		provider := NewJSONProvider(file.NewPath(os.Args[0], "/service_config/service.json"), config)
		provider.Provide()
	}
	return config
}

//ServiceConfiguration ...
type ServiceConfiguration struct {
	ExternalAPIHost       string `json:"ExternalAPIHost"`
	MySQLConnectionString string `json:"MySQLConnectionString"`
	ServiceHost           string `json:"ServiceHost"`
	ServicePort           int    `json:"ServicePort"`
	SchedulerDuration     int    `json:"SchedulerDurationSec"`
}
