package middlewares

import "github.com/gin-gonic/gin"

func AuthMiddleware(c *gin.Context) {
	//token := c.GetHeader("Authorization")
	//
	c.Next()
}
