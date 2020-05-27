package middleware

import (
	"github.com/gin-gonic/gin"
)

// GetMiddleware preps the middleware as advised by the config
func GetMiddleware(router **gin.Engine, url string) gin.IRoutes {
	var v1 gin.IRoutes

	v1 = (*router).Group(url)
	return v1
}
