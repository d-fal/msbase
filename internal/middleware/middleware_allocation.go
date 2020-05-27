package middleware

import (
	"errors"
	"fmt"
	"msbase/pkg/logger"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/logrusorgru/aurora"
)

// AllocateMiddleware allcocates middleware

func AllocateMiddleware(mwID string, id string, routes *gin.IRoutes) error {

	switch mwID {
	case middlewaresSet.CORSMiddleware:
		(*routes).Use(cors.Default())

	case middlewaresSet.Authenticator:
		(*routes).Use(Authentication)

	case middlewaresSet.LoggerMiddleware:
		(*routes).Use(LoggerMiddleware(logger.GetZapLogger(id), time.RFC3339, true))

	default:
		return errors.New(fmt.Sprintf("App profile error. "+
			"The requested middleware {%s} does not exist. "+
			"Please make sure there are no mistakes in your identity file."+
			"Currently, ther available middlewares are: %v ", aurora.Yellow(mwID), aurora.Green(middlewaresSet)))
	}
	return nil
}
