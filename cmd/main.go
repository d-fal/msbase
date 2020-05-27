package main

import (
	"crypto/tls"
	"fmt"
	"msbase/pkg/conf"
	"msbase/pkg/identity"
	"msbase/pkg/logger"
	"msbase/pkg/model"
	"syscall"

	"net/http"
	"os"
	"os/signal"

	"msbase/pkg/utils"

	"msbase/internal/loader"

	"github.com/gin-gonic/gin"
)

var (
	basePath string
)

//go:generate go run ../gen/main.go
func init() {
	model.CreateModels()
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	basePath = "config"
	gin.SetMode(gin.DebugMode)

}

func main() {
	c := make(chan os.Signal)

	signal.Notify(c, os.Interrupt)
	signal.Notify(c, syscall.SIGHUP)

	go func() {
		select {
		case sig := <-c:
			fmt.Println("Interrupt ", sig)
			zapLogger := logger.GetZapLogger(identity.GetSignature().App.ID)
			zapLogger.Sugar().Info("App stopped")
			utils.Terminate()
			os.Exit(1)
		}
	}()
	if len(os.Args[1:]) >= 1 {
		fmt.Println("Overriding logpath")
		basePath = os.Args[1]
	}
	configObject := conf.GetConfigObject()
	configObject.SetBasePath(basePath)
	configObject.Init()

	loader.LoadPackages()
}
