package conf

// Cache determines which cache server is used
type Cache struct {
	Name      string `yaml:"Name"`
	MaxIdle   int    `yaml:"MaxIdle"`
	MaxActive int    `yaml:"MaxActive"`
	Proto     string `yaml:"Proto"`
	Address   string `yaml:"Address"`
}

// GetCacheConfig flowbrings up the cache
func (confRcv *ConfigRcv) GetCacheConfig() Cache {
	return ConfigList.Cache
}
