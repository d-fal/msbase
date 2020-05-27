package conf

// ServerConfig config of the server
type ServerConfig struct {
	HTTPPort     string `yaml:"HTTPPort"`
	HTTPSPort    string `yaml:"HTTPSPort"`
	BasePath     string `yaml:"BasePath"`
	AssetsFolder string `yaml:"AssetsFolder"`
	IconsFolder  string `yaml:"IconsFolder"`
	Static       struct {
		URL  string `yaml:"URL"`
		Path string `yaml:"Path"`
	} `yaml:"Static"`
	BaseURL        string `yaml:"BaseURL"`
	RequestTimeout int    `yaml:"RequestTimeout"`
	LogFile        string `yaml:"LogFile"`
	AuthServerID   string `yaml:"AuthServerID"`
}

// GetServerConfig delivers the server config
func (confRcv *ConfigRcv) GetServerConfig() ServerConfig {
	return ConfigList.ServerConfig
}
