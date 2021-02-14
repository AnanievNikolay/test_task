package configuration

import (
	"strings"

	"github.com/AnanievNikolay/test_task/app/file"
)

//NewSettings ...
func NewSettings(_path *file.Path) *ServiceSettings {
	settings := &ServiceSettings{}
	//intializer := NewYAMLProvider(file.NewPath(os.Args[0], "/service_config/settings.yaml"), settings)
	intializer := NewYAMLProvider(_path, settings)
	intializer.Provide()
	return settings
}

//ServiceSettings ...
type ServiceSettings struct {
	Fsym []string `yaml:"fsyms"`
	Tsym []string `yaml:"tsyms"`
}

//Fsyms ...
func (setting *ServiceSettings) Fsyms() string {
	return strings.ToUpper(strings.Join(setting.Fsym, ","))
}

//Tsyms ..
func (setting *ServiceSettings) Tsyms() string {
	return strings.ToUpper(strings.Join(setting.Tsym, ","))
}
