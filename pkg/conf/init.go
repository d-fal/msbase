package conf

import (
	"fmt"
)

func init() {

	confRcv = &ConfigRcv{}
	confRcv.SetBasePath("config")
	confRcv.Init()
}

// Init initiates config readers
func (confRcv *ConfigRcv) Init() {

	LoadConfigFile(fmt.Sprintf("%s/errors.yaml", configPath), &errorMap)
	LoadConfigFile(fmt.Sprintf("%s/app_params.yaml", configPath), &ConfigList)
	LoadConfigFile(fmt.Sprintf("%s/lexicon.yaml", configPath), &lexicon)

}

func GetConfigObject() *ConfigRcv {

	return confRcv
}
