package conf

import (
	"io/ioutil"
	"log"

	"github.com/logrusorgru/aurora"
	"gopkg.in/yaml.v2"
)

// LoadConfigFile prepares config
func LoadConfigFile(filePath string, targetStruct interface{}) error {

	yamlFile, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Printf("Error: yaml.Get err   #%v , file does not exist\n%s\n", err, aurora.Yellow(filePath))
		return err
	}
	err = yaml.Unmarshal(yamlFile, targetStruct)
	if err != nil {
		return err
	}

	return err

}
