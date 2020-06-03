
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
