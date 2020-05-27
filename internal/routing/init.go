package routing

import (
	"github.com/gin-gonic/gin"
)

type HandlersSet struct {
	/* example extra handlers*/

}

func init() {

	handlersMap = make(map[string]gin.HandlerFunc)

}
