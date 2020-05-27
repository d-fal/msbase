package routing

import "github.com/gin-gonic/gin"

var (
	handlersSet HandlersSet
	handlersMap map[string]gin.HandlerFunc
)

const (
	MicrsoserviceCallMethodPOST = "POST"
	MicrsoserviceCallMethodGET  = "GET"
)
