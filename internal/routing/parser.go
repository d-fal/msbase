package routing

import (
	"fmt"
	"log"
	"regexp"

	"msbase/pkg/conf"
	"msbase/pkg/identity"

	"github.com/logrusorgru/aurora"
)

// ParseSignature parses the signature
func ParseSignature() identity.Parameters {

	if conf.LoadConfigFile(fmt.Sprintf("%s/app_params.yaml", conf.GetConfigObject().GetConfigPath()), &signatureObject) != nil {
		log.Fatalf("Malformed file config. Please check if %s exists. This file should be ", aurora.Red("app_paramss.yaml"))
	}
	identity.SetSignature(signatureObject)
	signatureObject = identity.GetSignature()

	return signatureObject
}

// SetVersionIntoURL this func stamp the version to its correponding url
func SetVersionIntoURL(serviceUrl string, version float64) string {

	pattern := regexp.MustCompile("{{.version}}")
	return pattern.ReplaceAllString(serviceUrl, fmt.Sprintf("v%d", int(version)))
}
