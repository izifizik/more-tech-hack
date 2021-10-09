package middlewares

import (
	"context"
	"github.com/gin-gonic/gin"
	"more-tech-hack/internal/config"
	"net/http"
)

func AuthMiddleware(c *gin.Context) {
	_, _, err := config.Client.DecodeAccessToken(context.Background(), c.GetHeader("Authorization"), config.KeyRealm,"")
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": "invalid token " + err.Error(),
		})
		return
	}
	c.Next()
}
