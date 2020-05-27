package utils

import (
	"reflect"

	"github.com/gin-gonic/gin"

	"msbase/pkg/conf"
	"msbase/pkg/libs"
)

// PrepareThirdPartyResponseDescription prepares data block that is the currency to this system messages
func PrepareThirdPartyResponseDescription(desc conf.MessageMicroBlock, label string, data interface{}) interface{} {
	v := reflect.ValueOf(desc)
	resp := make(map[string]interface{})
	types := v.Type()

	for i := 0; i < v.NumField(); i++ {
		resp[types.Field(i).Name] = v.Field(i).Interface()
	}
	resp[libs.BlockTagLabel] = label
	resp[libs.BlockDataLabel] = data
	return resp
}

// prepareOutput this func creates and prepares output response
func PrepareOutput(response interface{}, errorBlock *conf.ErrorBlock) interface{} {

	var responseArray interface{}

	if (*errorBlock).IsOk() {

		switch reflect.TypeOf(response).Kind() {

		case reflect.Map:

			responseArray = serializeResponse(response, errorBlock)

		case reflect.Slice:

			responseArray = response

		}

	} else {
		(*errorBlock) = conf.GetConfigObject().GetErrorList().FailedToParseServerResponse
	}
	return responseArray

}

func serializeResponse(data interface{}, errorBlock *conf.ErrorBlock) []interface{} {
	var (
		response []interface{}
	)

	if data != nil {
		switch reflect.TypeOf(data).Kind() {
		case reflect.Map:

			if len(data.(map[string]interface{})) > 0 {
				for _, v := range data.(map[string]interface{}) {
					response = append(response, v)
				}
			}
		}
	}
	return response
}

func ReflectServerOutput(c *gin.Context, response interface{}, errorBlock *conf.ErrorBlock) {

	if (*errorBlock).IsOk() {
		(*errorBlock) = conf.GetConfigObject().GetErrorList().UnknownError
	}

	errorReponse, errorCode, ok := GetHTTPStatus(errorBlock)

	c.JSON(
		errorCode,
		gin.H{
			"error": errorReponse,
			"data":  response,
		},
	)
	if ok {
		c.Abort()
	}
}
