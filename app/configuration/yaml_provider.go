package configuration

import (
	"io/ioutil"
	"log"

	"github.com/AnanievNikolay/test_task/app/file"
	"gopkg.in/yaml.v2"
)

//NewYAMLProvider ...
func NewYAMLProvider(_filePath *file.Path, _object interface{}) *YAMLProvider {
	return &YAMLProvider{
		filePath: _filePath,
		object:   _object,
	}
}

//YAMLProvider ...
type YAMLProvider struct {
	filePath *file.Path
	object   interface{}
}

//Provide ...
func (provider *YAMLProvider) Provide() {
	yamlFile, err := ioutil.ReadFile(provider.filePath.Abs())
	if err != nil {
		log.Fatalln("Error while get yaml file. Error:", err.Error())
	}
	err = yaml.Unmarshal(yamlFile, provider.object)
	if err != nil {
		log.Fatalln("Error while unmarshaling yaml file. Error: ", err.Error())
	}
}
