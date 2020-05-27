package routing

import (
	"fmt"

	"msbase/internal/middleware"
	"msbase/pkg/conf"
	"msbase/pkg/identity"

	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
)

var (
	signatureObject identity.Parameters
)

// Router sets up the router
func Router() (*gin.Engine, error) {
	signatureObject = ParseSignature()
	router := SetupRouter()

	router.Use(static.Serve("/static",
		static.LocalFile(fmt.Sprintf("%s",
			conf.GetConfigObject().GetServerConfig().Static.Path), false)))

	signatureObject.App.BaseURL = SetVersionIntoURL(signatureObject.App.BaseURL,
		signatureObject.App.Version)

	for _, urlObject := range signatureObject.App.Routes {
		v1 := middleware.GetMiddleware(&router, signatureObject.App.BaseURL)

		requestHandler, errorHandler := getRequestHandler(urlObject.ID)
		if errorHandler != nil {
			return router, errorHandler
		}

		for _, mw := range urlObject.Middlewares {
			if err := middleware.AllocateMiddleware(mw.ID,
				fmt.Sprintf("%s%s", signatureObject.App.BaseURL, urlObject.URL),
				&v1); err != nil {
				return nil, err
			}
		}

		switch urlObject.Method {

		case MicrsoserviceCallMethodPOST:
			v1.POST(urlObject.URL, requestHandler)

		case MicrsoserviceCallMethodGET:
			v1.GET(urlObject.URL, requestHandler)
		}

	}

	return router, nil
}

// SetupRouter sets up the router
func SetupRouter() *gin.Engine {
	return gin.Default()
}
