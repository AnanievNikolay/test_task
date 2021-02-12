package configuration

import (
	"log"

	"github.com/AnanievNikolay/test_task/app/file"
	"github.com/tkanos/gonfig"
)

//NewJSONProvider ...
func NewJSONProvider(_filePath *file.Path, _object interface{}) IProvider {
	return &JSONProvider{
		filePath: _filePath,
		object:   _object,
	}
}

//JSONProvider ...
type JSONProvider struct {
	filePath *file.Path
	object   interface{}
}

//Provide ...
func (provider *JSONProvider) Provide() {
	log.Println("Loading service configuration from:", provider.filePath.Abs())
	err := gonfig.GetConf(provider.filePath.Abs(), provider.object)
	if err != nil {
		log.Fatal("Unable to load configuration from:" + provider.filePath.Abs())
	}
}
