package middleware

import (
	"fmt"
	"msbase/pkg/utils"

	"github.com/gin-gonic/gin"
	"github.com/logrusorgru/aurora"
)

func Authentication(c *gin.Context) {
	token := c.Request.Header.Get("Authorization")
	token = token[len("bearer "):]

	_, errorBlock := authServer.IntrospectToken(token)

	if errorBlock.IsOk() {
		fmt.Println("TOKEN: ", aurora.Yellow(errorBlock))
		c.Set("token", token)
		c.Next()
	} else {
		utils.ReflectServerOutput(c, nil, &errorBlock)
		c.Abort()
	}

}
