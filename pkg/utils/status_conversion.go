package utils

import (
	"net/http"
	"reflect"

	"msbase/pkg/conf"

	"msbase/pkg/libs"
)

// GetHTTPStatus converst status to http code
func GetHTTPStatus(customError *conf.ErrorBlock) (interface{}, int, bool) {
	errorCode := 0

	if ((*customError).Code == nil && (*customError).HTTPCode == nil) || customError.IsOk() {
		(*customError).Code = conf.GetConfigObject().GetErrorList().NoError.Code
		(*customError).Message = conf.GetConfigObject().GetErrorList().NoError.Message
		(*customError).MessageFarsi = conf.GetConfigObject().GetErrorList().NoError.MessageFarsi
	}

	shouldTerminate := false

	if customError.HTTPCode == nil {
		switch customError.Status {
		case libs.StatusSuccess:
			errorCode = http.StatusOK
		case libs.StatusFailed:
			errorCode = http.StatusInternalServerError
			shouldTerminate = true
		case libs.StatusError:
			errorCode = http.StatusNotFound
			shouldTerminate = true
		default:
			errorCode = http.StatusOK
		}
	} else {
		switch customError.HTTPCode.(type) {
		case int:
			errorCode = customError.HTTPCode.(int)
			shouldTerminate = true
		}
	}

	e := reflect.ValueOf(customError).Elem()

	errorResponse := make(map[string]interface{})

	for i := 0; i < e.NumField(); i++ {
		fieldName := e.Type().Field(i).Name
		if fieldName == "HTTPCode" {
			continue
		}
		errorResponse[fieldName] = e.Field(i).Interface()
	}
	return errorResponse, errorCode, shouldTerminate
}
