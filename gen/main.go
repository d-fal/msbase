package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"text/template"

	"github.com/logrusorgru/aurora"
)

type Data struct {
	PackageName string
	HandlerID   string
}

type Activator struct {
	Imports     string
	Activations string
}

func createFromTemplate(path string, tmpl string, params interface{}) {
	file, err := os.Create(path)
	if err != nil {
		log.Fatal("Cannot create/update  ", path, err)
	}
	tmp, err := template.New("microservice").Parse(tmpl)

	if err != nil {
		fmt.Println("Cannot parse template ", err)
	}
	tmp.Execute(file, params)
}
func main() {

	fmt.Printf("\n\tMicroservice generator wizard\n\t\tWelcome!\n\t\t%s\n\t\t%s\n\n"+
		"You are about to generate your own "+
		"microservice. To start with, why not creating your microservice under \n"+
		"\t\t %s ?\nIf you have already done it, here we go....\n-----------------------------------\n\n",
		aurora.Yellow("Microservice generator wizard is a tool to seamlessly create microservice structures and basics for "+
			"developers."), aurora.Green("Developers are invited to browse README and drop their issues there."),
		aurora.Blue("<PROJECT_ROOT>/internal/microservice/<YOUR_OWN_MICROSERVICE>"))

	microservicePath := "../internal/microservice"
	files, err := ioutil.ReadDir(microservicePath)

	if err != nil {
		log.Fatal(err)
	}
	activator := Activator{}

	for _, f := range files {

		_, err := os.Stat(fmt.Sprintf("%[1]v/%[2]v/%[2]v.go", microservicePath, f.Name()))
		if err == nil {
			activator.Imports += fmt.Sprintf("\t\"msbase/internal/microservice/%s\"\n", f.Name())
			activator.Activations += fmt.Sprintf("\t%s.ActivateMicroservice()\n", f.Name())
			initializePackage(microservicePath, f.Name())
		}
		if os.IsNotExist(err) {
			fmt.Printf("%s is a new microservice, we are about to generate code for it\n", aurora.Red(f.Name()))
			if err := os.Mkdir(fmt.Sprintf("%s/%s/services", microservicePath, f.Name()), os.ModePerm); err != nil {
				fmt.Println("Services folder already exists")
			}
			createFromTemplate(fmt.Sprintf("%s/%s/services/services.go", microservicePath, f.Name()), newServicePackageTemplate, nil)
			createFromTemplate(fmt.Sprintf("%s/%s/services/handlers.go", microservicePath, f.Name()), newServiceHandlerFile, nil)
			initializePackage(microservicePath, f.Name())
		}
	}
	updateMicroservices(activator)
	// createRegistrar(activator)
	//go:generate go run generate_loaders.go

}

func updateMicroservices(activator Activator) {
	file, err := os.Create("../internal/microservice/microservice.go")
	if err != nil {
		log.Fatal("Cannot create/update microservice.go ", err)
	}
	tmp, err := template.New("microservice").Parse(microserviceTemplate)

	if err != nil {
		fmt.Println("Cannot parse template ", err)
	}
	tmp.Execute(file, activator)

}
func updateActivator(activator Data) {
	file, err := os.Create(fmt.Sprintf("../internal/microservice/%[1]v/%[1]v.go", activator.PackageName))
	if err != nil {
		log.Fatalf("Cannot create/update %s.go %v\n", activator.PackageName, err)
	}
	tmp, err := template.New("microservice").Parse(activatorTemplate)

	if err != nil {
		fmt.Println("Cannot parse template ", err)
	}
	tmp.Execute(file, activator)
}

func initializePackage(pPath, pName string) {
	newPackage := Data{}
	newPackage.PackageName = pName

	t, _ := template.New("package").Parse(packageTemplate)
	file, err := os.Create(fmt.Sprintf("%[1]v/%[2]v/%[2]v.go", pPath, pName))
	if err != nil {
		log.Fatal("Cannot create file", err)
	}
	if err = t.Execute(file, newPackage); err != nil {
		log.Fatal("Cannot assign template", err)
	}
	updateActivator(newPackage)

}

var (
	packageTemplate = `
package {{.PackageName}} 

import (
	"msbase/internal/microservice/{{.PackageName}}/services"
	"msbase/internal/routing"
)
var (
	handler *services.ServiceHandler
)

func GetHandler() *services.ServiceHandler {
	return handler
}
func ActivateMicroservice() {
	handler = services.GetHandler()
	// routing.RegisterHandler(handler.GinHandlerName, "internal_rec_get_auth_token")
}
`
	microserviceTemplate = `
package microservice

import (
{{.Imports}}
)

// LoadActiveMicroservices this method in close connection to
/*
***********************************************************
*  Note: this function should be used to activate the
*
***********************************************************


 */
func LoadActiveMicroservices() {

{{.Activations}}
}
`

	activatorTemplate = `
package {{.PackageName}}

import (
	"msbase/internal/routing"
	"reflect"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/logrusorgru/aurora"
	"fmt"
	"msbase/internal/microservice/{{.PackageName}}/services"
)

func ActivateMicroservice() {
	handler := services.GetHandler()

	handlerType := reflect.ValueOf(handler)
	for i := 0; i < handlerType.NumMethod(); i++ {
		method := handlerType.Method(i)
		_handler := method.Interface().(func(*gin.Context))
		_handlerID := fmt.Sprintf("internal_rec_{{.PackageName}}_%s", strings.ToLower(reflect.TypeOf(handler).Method(i).Name))
		fmt.Println("Given handlerID: Use it in your app_params.yaml ", aurora.Cyan(_handlerID))
		routing.RegisterHandler(_handler, _handlerID)

	}
}

`
	newServicePackageTemplate = `
package services

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

/* GinHandlerName is the name of handler you want in this project
GIN Handlers Shuould Be Added Here
@Note: This should be introduced to the project in internal/routing/handlers.go
*/
func (handler *ServiceHandler) GinTestHandler(c *gin.Context) {
	c.JSON(
		http.StatusOK,
		gin.H{
			"result": "OK",
		},
	)
}
`
	newServiceHandlerFile = `
package services

// ServiceHandler handling the base functionality
type ServiceHandler struct {
	Name string
}

var (
	handler *ServiceHandler
)

func GetHandler() *ServiceHandler {
	return handler
}

`
)
