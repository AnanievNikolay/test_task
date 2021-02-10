package configuration

var config *ServiceConfiguration

//ServiceConfig ...
func ServiceConfig() *ServiceConfiguration {
	if config == nil {
	}
	return config
}

//ServiceConfiguration ...
type ServiceConfiguration struct {
	ExternalHost          string `json:"ExternalHost"`
	MySQLConnectionString string `json:"MySQLConnectionString"`
	ServiceHost           string `json:"ServiceHost"`
	ServicePort           int    `json:"ServicePort"`
	ExternalAPIHost       string `json:"ExternalAPIHost"`
}
