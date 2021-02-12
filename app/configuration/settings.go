package configuration

import (
	"os"

	"github.com/AnanievNikolay/test_task/app/file"
)

var settings *ServiceSettings

//Settings ...
func Settings() *ServiceSettings {
	if settings == nil {
		settings = &ServiceSettings{}
		intializer := NewYAMLProvider(file.NewPath(os.Args[0], "/service_config/settings.yaml"), settings)
		intializer.Provide()
	}
	return settings
}

//ServiceSettings ...
type ServiceSettings struct {
	Fsym []string `yaml:"fsyms"`
	Tsym []string `yaml:"tsyms"`
}
