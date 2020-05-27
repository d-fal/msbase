package conf

// GetDatabaseConfig generates database config
func (confRcv *ConfigRcv) GetDatabaseConfig() DatabaseObject {
	return ConfigList.Database
}
