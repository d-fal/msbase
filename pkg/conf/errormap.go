package conf

// SetLogPath set path of the log file
func SetLogPath(path string) {
	configPath = path
}

// GetErrorList load the list of errors
func (confRcv *ConfigRcv) GetErrorList() ErrorTypes {
	return errorMap.Errors
}

// GetCostumError prepares the costum error as of defined in the config
func (confRcv *ConfigRcv) GetCostumError(errCode interface{}, errorBlocks []ErrorBlock) ErrorBlock {
	errorCode := 0
	switch errCode.(type) {
	case int:
		errorCode = errCode.(int)
	case float64:
		errorCode = int(errCode.(float64))
	}
	for _, errBlck := range errorBlocks {
		if errBlck.Code == errorCode {
			return errBlck
		}
	}
	return ErrorBlock{}
}
