
package test

import (
	"msbase/internal/routing"
	"reflect"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/logrusorgru/aurora"
	"fmt"
	"msbase/internal/microservice/test/services"
)

func ActivateMicroservice() {
	handler := services.GetHandler()

	handlerType := reflect.ValueOf(handler)
	for i := 0; i < handlerType.NumMethod(); i++ {
		method := handlerType.Method(i)
		_handler := method.Interface().(func(*gin.Context))
		_handlerID := fmt.Sprintf("internal_rec_test_%s", strings.ToLower(reflect.TypeOf(handler).Method(i).Name))
		fmt.Println("Given handlerID: Use it in your app_params.yaml ", aurora.Cyan(_handlerID))
		routing.RegisterHandler(_handler, _handlerID)

	}
}

