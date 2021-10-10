package middlewares

import (
	"context"
	"github.com/gin-gonic/gin"
	"log"
	"more-tech-hack/internal/config"
	"net/http"
)

func AuthMiddleware(c *gin.Context) {
	_, _, err := config.Client.DecodeAccessToken(context.Background(), c.GetHeader("Authorization"), config.KeyRealm,"")
	if err != nil {
		log.Println(err)
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": "invalid token " + err.Error(),
		})
		return
	}
	c.Next()
}
