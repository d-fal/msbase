package middleware

import "github.com/gin-gonic/gin"

func BaseMiddleware(c *gin.Context) {
	c.Next()
}
