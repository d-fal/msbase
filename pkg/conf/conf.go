package conf

import (
	"errors"
	"regexp"
	"strconv"
)

var (
	r          *regexp.Regexp
	trim       *regexp.Regexp
	configPath string
	ConfigList Config
	confRcv    *ConfigRcv
)

type ConfigRcv struct {
	Name string
}

// Config gets all the defined errors
type Config struct {
	ServerConfig    ServerConfig   `yaml:"ServerConfig"`
	EventRecipients EventRecipient `yaml:"EventRecipient"`
	Database        DatabaseObject `yaml:"Database"`
	Cache           Cache          `yaml:"Cache"`
}

// EventRecipient event recipients that handle the config file
type EventRecipient struct {
	URL            string `yaml:"url"`
	Script         string `yaml:"script"`
	LogFile        string `yaml:"logFile"`
	LogPath        string `yaml:"logPath"`
	MinRespiteTime int64  `yaml:"minRespiteTime"` // Minimum time (in seconds) to wait for sending next alert
	ErrorThreshold int64  `yaml:"errorThreshold"` // Minimum errors to get before sending the report
	SearchedPhrase string `yaml:"searchedPattern"`
	ReportTo       string `yaml:"reportTo"`
}

// DatabaseObject the config parameters that define the database connection
type DatabaseObject struct {
	Username                  string `yaml:"Username"`
	Password                  string `yaml:"Password"`
	SID                       string `yaml:"SID"`
	MinSessions               int    `yaml:"MinSessions"`
	MaxSessions               int    `yaml:"MaxSessions"`
	PoolIncrement             int    `yaml:"PoolIncrement"`
	ConnClass                 string `yaml:"ConnClass"`
	ConnectionMaximumLifeTime int    `yaml:"ConnectionMaximumLifeTime"`
	DBConnectMaximumAttempts  int    `yaml:"DBConnectMaximumAttempts"`
}

// SetBasePath sets the path to config files
func (confRcv *ConfigRcv) SetBasePath(path string) {

	confRcv.Name = path
	configPath = path
}

//GetConfigPath returns config path
func (confRcv *ConfigRcv) GetConfigPath() string {
	return configPath
}

// FindPattern expects a pointer to search string an returns expected array
func (confRcv *ConfigRcv) FindPattern(searchStr *string) (string, int, error) {

	var (
		url        string
		statusCode int
		err        error
	)
	result := r.FindAllString(*searchStr, -1)

	if len(result) < 2 {
		return url, statusCode, errors.New(confRcv.GetErrorList().ErrorPatternNotFound.Message)
	}
	/* Remove tailing or heading whitespaces */
	for i := 0; i < len(result); i++ {
		result[i] = trim.ReplaceAllString(result[i], "")
	}
	if statusCode, err = strconv.Atoi(result[0]); err != nil {
		return url, statusCode, errors.New(confRcv.GetErrorList().ErrorStatusCodeNotFound.Message)
	}
	/* This is safe, the above code prevent checking this line if its length is lower than 2*/
	url = result[1]
	if url == "" {
		return url, statusCode, errors.New(confRcv.GetErrorList().ErrorURLNotFound.Message)
	}
	return url, statusCode, nil /* error is nil when there is any*/
}
